// Copyright 2020-present Open Networking Foundation.
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
	devicechange "github.com/onosproject/onos-config/api/types/change/device"
	"github.com/onosproject/onos-config/api/types/device"
	netchangestore "github.com/onosproject/onos-config/pkg/store/change/network"
	devicesnapstore "github.com/onosproject/onos-config/pkg/store/snapshot/device"
)

// newDeviceStateManager returns a new device state manager
func newDeviceStateManager(changeStore netchangestore.Store, snapshotStore devicesnapstore.Store) *deviceStateManager {
	return &deviceStateManager{
		changeStore:   changeStore,
		snapshotStore: snapshotStore,
		devices:       make(map[device.VersionedID]*deviceState),
	}
}

// deviceStateManager is a registry of device states
type deviceStateManager struct {
	changeStore   netchangestore.Store
	snapshotStore devicesnapstore.Store
	devices       map[device.VersionedID]*deviceState
}

// getDeviceState returns the state for the given device
func (m *deviceStateManager) getDeviceState(deviceID device.VersionedID) (*deviceState, error) {
	device, ok := m.devices[deviceID]
	if !ok {
		state := make(map[string]devicechange.Index)
		snapshot, err := m.snapshotStore.Load(deviceID)
		if err != nil {
			return nil, err
		} else if snapshot != nil {
			for _, value := range snapshot.Values {
				state[value.Path] = snapshot.ChangeIndex
			}
		}

		device = &deviceState{
			state: state,
		}
		m.devices[deviceID] = device
	}
	return device, nil
}

// deviceState manages the state for a single device
type deviceState struct {
	state map[string]devicechange.Index
}

// get gets the index of the given path
func (s *deviceState) get(path string) devicechange.Index {
	return s.state[path]
}

// update updates the index of the given path
func (s *deviceState) update(path string, index devicechange.Index) {
	s.state[path] = index
}
