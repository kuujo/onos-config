package request

import (
	"github.com/onosproject/onos-config/pkg/controller"
	"github.com/onosproject/onos-config/pkg/controller/network"
	leadershipstore "github.com/onosproject/onos-config/pkg/store/leadership"
	networkstore "github.com/onosproject/onos-config/pkg/store/network"
	requeststore "github.com/onosproject/onos-config/pkg/store/request"
	networktype "github.com/onosproject/onos-config/pkg/types/network"
	requesttype "github.com/onosproject/onos-config/pkg/types/request"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
)

// NewController returns a new network controller
func NewController(leadership leadershipstore.Store, requests requeststore.Store, networks networkstore.Store) *controller.Controller {
	c := controller.NewController()
	c.Activate(&controller.LeadershipActivator{
		Store: leadership,
	})
	c.Watch(&Watcher{
		Store: requests,
	})
	c.Watch(&network.OwnerWatcher{
		Store: networks,
	})
	c.Reconcile(&Reconciler{
		requests: requests,
		networks: networks,
	})
	return c
}

// Reconciler is the request reconciler
type Reconciler struct {
	requests requeststore.Store
	networks networkstore.Store
	lowMark  uint64
}

func (r *Reconciler) Reconcile(id interface{}) (bool, error) {
	return r.reconcile(id.(requesttype.ID))
}

func (r *Reconciler) reconcile(id requesttype.ID) (bool, error) {
	request, err := r.requests.Get(id)
	if err != nil {
		return false, err
	}

	// Check if a network configuration has been created.
	network, err := r.networks.Get(request.GetNetworkID())
	if err != nil {
		return false, nil
	}

	// If the network has not been created, add it in the PENDING state
	if network == nil {
		network = &networktype.NetworkConfig{
			ID:        request.GetNetworkID(),
			RequestID: request.ID,
			Status:    networktype.Status_PENDING,
			Changes:   request.Changes,
		}
		err = r.networks.Create(network)
		if err != nil {
			return false, nil
		}
	}

	// Replay requests starting at the low water mark.
	ch := make(chan *requesttype.ConfigRequest)
	if err := r.requests.Replay(requesttype.ID(r.lowMark), ch); err != nil {
		return false, err
	}

	// Loop through all requests that exist prior to the given request. For any request that is still pending,
	// if the previous configuration's devices intersect with the current configuration's devices then requeue
	// the request for reconciliation. We must wait for prior configurations to be applied before we can continue.
	allComplete := true
	for prevRequest := range ch {
		if prevRequest.ID < request.ID {
			network, err := r.networks.Get(prevRequest.GetNetworkID())
			if err != nil {
				return false, err
			}

			// If all prior requests have been completed and this request is completed, increment the low water mark.
			if allComplete && network.Status == networktype.Status_SUCCEEDED || network.Status == networktype.Status_FAILED {
				r.lowMark++
			} else {
				allComplete = false
			}

			// If the previous request's devices intersect with this request's devices and the previous request
			// is not yet complete, requeue the request.
			if isIntersecting(request.GetDevices(), prevRequest.GetDevices()) && network.Status != networktype.Status_SUCCEEDED && network.Status != networktype.Status_FAILED {
				return false, nil
			}
		}
	}

	// If we've made it this far, update the network config state to APPLYING.
	network.Status = networktype.Status_APPLYING
	err = r.networks.Update(network)
	if err != nil {
		return false, err
	}
	return true, nil
}

var _ controller.Reconciler = &Reconciler{}

// isIntersecting returns a bool indicating whether the given sets of devices intersect
func isIntersecting(devices1, devices2 []device.ID) bool {
	for _, device1 := range devices1 {
		for _, device2 := range devices2 {
			if device1 == device2 {
				return true
			}
		}
	}
	return false
}
