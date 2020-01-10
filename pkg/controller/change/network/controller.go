// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	"github.com/onosproject/onos-config/api/types"
	changetypes "github.com/onosproject/onos-config/api/types/change"
	devicechange "github.com/onosproject/onos-config/api/types/change/device"
	networkchange "github.com/onosproject/onos-config/api/types/change/network"
	"github.com/onosproject/onos-config/api/types/device"
	"github.com/onosproject/onos-config/pkg/controller"
	devicechangestore "github.com/onosproject/onos-config/pkg/store/change/device"
	networkchangestore "github.com/onosproject/onos-config/pkg/store/change/network"
	devicestore "github.com/onosproject/onos-config/pkg/store/device"
	"github.com/onosproject/onos-config/pkg/store/device/cache"
	leadershipstore "github.com/onosproject/onos-config/pkg/store/leadership"
	devicesnapstore "github.com/onosproject/onos-config/pkg/store/snapshot/device"
	"github.com/onosproject/onos-config/pkg/utils/logging"
	devicetopo "github.com/onosproject/onos-topo/api/device"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var log = logging.GetLogger("controller", "change", "network")

// NewController returns a new config controller
func NewController(leadership leadershipstore.Store, deviceCache cache.Cache, devices devicestore.Store, networkChanges networkchangestore.Store, deviceChanges devicechangestore.Store, deviceSnapshots devicesnapstore.Store) *controller.Controller {
	c := controller.NewController("NetworkChange")
	c.Activate(&controller.LeadershipActivator{
		Store: leadership,
	})
	c.Watch(&Watcher{
		Store: networkChanges,
	})
	c.Watch(&DeviceWatcher{
		DeviceCache: deviceCache,
		DeviceStore: devices,
		ChangeStore: deviceChanges,
	})
	c.Reconcile(&Reconciler{
		networkChanges: networkChanges,
		deviceChanges:  deviceChanges,
		deviceStates:   newDeviceStateManager(networkChanges, deviceSnapshots),
		devices:        devices,
	})
	return c
}

// Reconciler is a config reconciler
type Reconciler struct {
	networkChanges networkchangestore.Store
	deviceChanges  devicechangestore.Store
	deviceStates   *deviceStateManager
	devices        devicestore.Store
}

// Reconcile reconciles the state of a network configuration
func (r *Reconciler) Reconcile(id types.ID) (bool, error) {
	change, err := r.networkChanges.Get(networkchange.ID(id))
	if err != nil {
		log.Warnf("Could not get NetworkChange %s", id)
		return false, err
	}

	log.Infof("Reconciling NetworkChange %v", change)

	// Handle the change for each phase
	if change != nil {
		switch change.Status.Phase {
		case changetypes.Phase_CHANGE:
			return r.reconcileChange(change)
		case changetypes.Phase_ROLLBACK:
			return r.reconcileRollback(change)
		}
	}
	return true, nil
}

// reconcileChange reconciles a change in the CHANGE phase
func (r *Reconciler) reconcileChange(change *networkchange.NetworkChange) (bool, error) {
	// If the change is not pending, skip it
	if change.Status.State != changetypes.State_PENDING {
		return true, nil
	}

	// Create device changes if necessary
	if !hasDeviceChanges(change) {
		return r.createDeviceChanges(change)
	}

	// Get the current state of all device changes for the change
	deviceChanges, err := r.getDeviceChanges(change)
	if err != nil {
		return false, err
	}

	// Ensure device changes are pending for the current incarnation
	changed, err := r.ensureDeviceChangesPending(change, deviceChanges)
	if changed || err != nil {
		return changed, err
	}

	// If the network change can be applied, apply it by incrementing the incarnation number
	apply, err := r.canTryChange(change, deviceChanges)
	if err != nil {
		return false, err
	} else if apply {
		change.Status.Incarnation++
		change.Status.State = changetypes.State_PENDING
		change.Status.Reason = changetypes.Reason_NONE
		change.Status.Message = ""
		log.Infof("Applying NetworkChange %v", change)
		if err := r.networkChanges.Update(change); err != nil {
			return false, err
		}
		return true, nil
	}

	// If all device changes are complete, complete the network change
	if r.isDeviceChangesComplete(change, deviceChanges) {
		change.Status.State = changetypes.State_COMPLETE
		log.Infof("Completing NetworkChange %v", change)
		if err := r.networkChanges.Update(change); err != nil {
			return false, err
		}
		return true, nil
	}

	// If any device change has failed, roll back all device changes
	if r.isDeviceChangesFailed(change, deviceChanges) {
		return r.ensureDeviceChangeRollbacks(change, deviceChanges)
	}
	return true, nil
}

// hasDeviceChanges indicates whether the given change has created device changes
func hasDeviceChanges(change *networkchange.NetworkChange) bool {
	return change.Refs != nil && len(change.Refs) > 0
}

// createDeviceChanges creates device changes in sequential order
func (r *Reconciler) createDeviceChanges(networkChange *networkchange.NetworkChange) (bool, error) {
	// If the previous network change has not created device changes, requeue to wait for changes to be propagated
	prevChange, err := r.networkChanges.GetByIndex(networkChange.Index - 1)
	if err != nil {
		return false, err
	} else if prevChange != nil && !hasDeviceChanges(prevChange) {
		return false, nil
	}

	// Loop through changes and create device changes
	refs := make([]*networkchange.DeviceChangeRef, len(networkChange.Changes))
	for i, change := range networkChange.Changes {
		state, err := r.deviceStates.getDeviceState(device.NewVersionedID(change.DeviceID, change.DeviceVersion))
		if err != nil {
			return false, err
		}

		// Record the index of the prior update for each path
		for _, value := range change.Values {
			value.PrevIndex = types.Index(state.get(value.Path))
		}

		deviceChange := &devicechange.DeviceChange{
			Index: devicechange.Index(networkChange.Index),
			NetworkChange: devicechange.NetworkChangeRef{
				ID:    types.ID(networkChange.ID),
				Index: types.Index(networkChange.Index),
			},
			Change: change,
		}
		log.Infof("Creating DeviceChange %v for %v", deviceChange, networkChange.ID)
		if err := r.deviceChanges.Create(deviceChange); err != nil {
			return false, err
		}

		// Record the index for each updated path
		for _, value := range change.Values {
			state.update(value.Path, deviceChange.Index)
		}

		refs[i] = &networkchange.DeviceChangeRef{
			DeviceChangeID: deviceChange.ID,
		}
	}

	// If references have been updated, store the refs and succeed the reconciliation
	networkChange.Refs = refs
	if err := r.networkChanges.Update(networkChange); err != nil {
		return false, err
	}
	return true, nil
}

// canTryChange returns a bool indicating whether the change can be attempted
func (r *Reconciler) canTryChange(change *networkchange.NetworkChange, deviceChanges []*devicechange.DeviceChange) (bool, error) {
	// If the incarnation number is positive, verify all device changes have been rolled back
	if change.Status.Incarnation > 0 {
		for _, deviceChange := range deviceChanges {
			if deviceChange.Status.Incarnation != change.Status.Incarnation ||
				deviceChange.Status.Phase != changetypes.Phase_ROLLBACK ||
				deviceChange.Status.State != changetypes.State_COMPLETE {
				return false, nil
			}
		}
	}

	// First, check if the devices affected by the change are available
	for _, deviceChange := range change.Changes {
		// Get the device from the topo service
		device, err := r.devices.Get(devicetopo.ID(deviceChange.DeviceID))
		if err != nil && status.Code(err) != codes.NotFound {
			return false, err
		} else if device == nil {
			return false, nil
		}

		// If the device state is not connected, skip this change
		state := getProtocolState(device)
		if state != devicetopo.ChannelState_CONNECTED {
			log.Infof("Cannot apply NetworkChange %v: %v is offline", change.ID, deviceChange.DeviceID)
			return false, nil
		}

		// For each change path/value, ensure the prior change is COMPLETE before processing this change
		for _, changeValue := range deviceChange.Values {
			if changeValue.PrevIndex == 0 {
				continue
			}

			// Get the prior change for the path
			prevChange, err := r.networkChanges.GetByIndex(networkchange.Index(changeValue.PrevIndex))
			if err != nil {
				return false, err
			}

			// If the change is not complete, do not process this change
			if prevChange.Status.State != changetypes.State_COMPLETE {
				return false, nil
			}
		}
	}
	return true, nil
}

// ensureDeviceChangesPending ensures device changes are pending
func (r *Reconciler) ensureDeviceChangesPending(networkChange *networkchange.NetworkChange, changes []*devicechange.DeviceChange) (bool, error) {
	// Ensure all device changes are being applied
	updated := false
	for _, deviceChange := range changes {
		if deviceChange.Status.Incarnation < networkChange.Status.Incarnation {
			deviceChange.Status.Incarnation = networkChange.Status.Incarnation
			deviceChange.Status.Phase = changetypes.Phase_CHANGE
			deviceChange.Status.State = changetypes.State_PENDING
			deviceChange.Status.Reason = changetypes.Reason_NONE
			log.Infof("Running DeviceChange %v", deviceChange)
			if err := r.deviceChanges.Update(deviceChange); err != nil {
				return false, err
			}
			updated = true
		}
	}
	return updated, nil
}

// getDeviceChanges gets the device changes for the given network change
func (r *Reconciler) getDeviceChanges(networkChange *networkchange.NetworkChange) ([]*devicechange.DeviceChange, error) {
	deviceChanges := make([]*devicechange.DeviceChange, len(networkChange.Changes))
	for i, changeRef := range networkChange.Refs {
		deviceChange, err := r.deviceChanges.Get(changeRef.DeviceChangeID)
		if err != nil {
			return nil, err
		}
		deviceChanges[i] = deviceChange
	}
	return deviceChanges, nil
}

// isDeviceChangesComplete checks whether the device changes are complete
func (r *Reconciler) isDeviceChangesComplete(networkChange *networkchange.NetworkChange, changes []*devicechange.DeviceChange) bool {
	for _, change := range changes {
		if change.Status.Incarnation != networkChange.Status.Incarnation ||
			change.Status.Phase != changetypes.Phase_CHANGE ||
			change.Status.State != changetypes.State_COMPLETE {
			return false
		}
	}
	return true
}

// isDeviceChangesFailed checks whether any device change has failed for the current incarnation
func (r *Reconciler) isDeviceChangesFailed(networkChange *networkchange.NetworkChange, changes []*devicechange.DeviceChange) bool {
	for _, change := range changes {
		if change.Status.Incarnation == networkChange.Status.Incarnation && change.Status.State == changetypes.State_FAILED {
			return true
		}
	}
	return false
}

// ensureDeviceChangeRollbacks ensures device changes are being rolled back
func (r *Reconciler) ensureDeviceChangeRollbacks(networkChange *networkchange.NetworkChange, changes []*devicechange.DeviceChange) (bool, error) {
	for _, deviceChange := range changes {
		if deviceChange.Status.Incarnation != networkChange.Status.Incarnation ||
			deviceChange.Status.Phase != changetypes.Phase_ROLLBACK ||
			deviceChange.Status.State == changetypes.State_FAILED {
			deviceChange.Status.Incarnation = networkChange.Status.Incarnation
			deviceChange.Status.Phase = changetypes.Phase_ROLLBACK
			deviceChange.Status.State = changetypes.State_PENDING
			log.Infof("Rolling back DeviceChange %v", deviceChange)
			if err := r.deviceChanges.Update(deviceChange); err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

// reconcileRollback reconciles a change in the ROLLBACK phase
func (r *Reconciler) reconcileRollback(change *networkchange.NetworkChange) (bool, error) {
	// Ensure the device changes are in the ROLLBACK phase
	updated, err := r.ensureDeviceRollbacks(change)
	if updated || err != nil {
		return updated, err
	}

	// If the change is not pending, skip it
	if change.Status.State != changetypes.State_PENDING {
		return true, nil
	}

	// Get the current state of all device changes for the change
	deviceChanges, err := r.getDeviceChanges(change)
	if err != nil {
		return false, err
	}

	// If the network rollback can be applied, apply it by incrementing the incarnation number
	apply, err := r.canTryRollback(change, deviceChanges)
	if err != nil {
		return false, err
	} else if apply {
		change.Status.Incarnation++
		change.Status.State = changetypes.State_PENDING
		change.Status.Reason = changetypes.Reason_NONE
		change.Status.Message = ""
		log.Infof("Rolling back NetworkChange %v", change)
		if err := r.networkChanges.Update(change); err != nil {
			return false, err
		}
		return true, nil
	}

	// If all device rollbacks are complete, complete the network change
	if r.isDeviceRollbacksComplete(change, deviceChanges) {
		change.Status.State = changetypes.State_COMPLETE
		log.Infof("Completing NetworkChange %v", change)
		if err := r.networkChanges.Update(change); err != nil {
			return false, err
		}
		return true, nil
	}

	// If any device change has failed, roll back all device changes
	if r.isDeviceChangesFailed(change, deviceChanges) {
		return r.ensureDeviceChangeRollbacks(change, deviceChanges)
	}
	return true, nil
}

// ensureDeviceRollbacks ensures device rollbacks are pending
func (r *Reconciler) ensureDeviceRollbacks(networkChange *networkchange.NetworkChange) (bool, error) {
	// Ensure all device changes are being rolled back
	updated := false
	for _, changeRef := range networkChange.Refs {
		deviceChange, err := r.deviceChanges.Get(changeRef.DeviceChangeID)
		if err != nil {
			return false, err
		}

		if deviceChange.Status.Incarnation < networkChange.Status.Incarnation ||
			deviceChange.Status.Phase != changetypes.Phase_ROLLBACK {
			deviceChange.Status.Incarnation = networkChange.Status.Incarnation
			deviceChange.Status.Phase = changetypes.Phase_ROLLBACK
			deviceChange.Status.State = changetypes.State_PENDING
			log.Infof("Rolling back DeviceChange %v", deviceChange)
			if err := r.deviceChanges.Update(deviceChange); err != nil {
				return false, err
			}
			updated = true
		}
	}
	return updated, nil
}

// isDeviceRollbacksComplete checks whether the device rollbacks are complete
func (r *Reconciler) isDeviceRollbacksComplete(networkChange *networkchange.NetworkChange, changes []*devicechange.DeviceChange) bool {
	for _, change := range changes {
		if change.Status.Incarnation != networkChange.Status.Incarnation ||
			change.Status.Phase != changetypes.Phase_ROLLBACK ||
			change.Status.State != changetypes.State_COMPLETE {
			return false
		}
	}
	return true
}

// canTryRollback returns a bool indicating whether the rollback can be attempted
func (r *Reconciler) canTryRollback(change *networkchange.NetworkChange, deviceChanges []*devicechange.DeviceChange) (bool, error) {
	// Verify all device changes are being rolled back
	if change.Status.Incarnation > 0 {
		for _, deviceChange := range deviceChanges {
			if deviceChange.Status.Incarnation != change.Status.Incarnation ||
				deviceChange.Status.Phase != changetypes.Phase_ROLLBACK ||
				deviceChange.Status.State != changetypes.State_FAILED {
				return false, nil
			}
		}
	}

	nextChange, err := r.networkChanges.GetNext(change.Index)
	if err != nil {
		return false, err
	}

	for nextChange != nil {
		// If the change intersects this change, verify it has been rolled back
		if isIntersectingChange(change, nextChange) {
			return nextChange.Status.Phase == changetypes.Phase_ROLLBACK &&
				(nextChange.Status.State == changetypes.State_COMPLETE ||
					nextChange.Status.State == changetypes.State_FAILED), nil
		}

		nextChange, err = r.networkChanges.GetNext(nextChange.Index)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// isIntersectingChange indicates whether the changes from the two given NetworkChanges intersect
func isIntersectingChange(config *networkchange.NetworkChange, history *networkchange.NetworkChange) bool {
	for _, configChange := range config.Changes {
		for _, historyChange := range history.Changes {
			if configChange.DeviceID == historyChange.DeviceID {
				return true
			}
		}
	}
	return false
}

func getProtocolState(device *devicetopo.Device) devicetopo.ChannelState {
	// Find the gNMI protocol state for the device
	var protocol *devicetopo.ProtocolState
	for _, p := range device.Protocols {
		if p.Protocol == devicetopo.Protocol_GNMI {
			protocol = p
			break
		}
	}
	if protocol == nil {
		return devicetopo.ChannelState_UNKNOWN_CHANNEL_STATE
	}
	return protocol.ChannelState
}

var _ controller.Reconciler = &Reconciler{}
