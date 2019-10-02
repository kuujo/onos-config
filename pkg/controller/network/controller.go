package network

import (
	"github.com/onosproject/onos-config/pkg/controller"
	"github.com/onosproject/onos-config/pkg/controller/change"
	changestore "github.com/onosproject/onos-config/pkg/store/change"
	leadershipstore "github.com/onosproject/onos-config/pkg/store/leadership"
	networkstore "github.com/onosproject/onos-config/pkg/store/network"
	changetype "github.com/onosproject/onos-config/pkg/types/change"
	networktype "github.com/onosproject/onos-config/pkg/types/network"
)

// NewController returns a new network controller
func NewController(leadership leadershipstore.Store, networks networkstore.Store, changes changestore.Store) *controller.Controller {
	c := controller.NewController()
	c.Activate(&controller.LeadershipActivator{
		Store: leadership,
	})
	c.Watch(&Watcher{
		Store: networks,
	})
	c.Watch(&change.OwnerWatcher{
		Store: changes,
	})
	c.Reconcile(&Reconciler{
		networks: networks,
		changes:  changes,
	})
	return c
}

// Reconciler is a network reconciler
type Reconciler struct {
	*controller.Controller
	networks networkstore.Store
	changes  changestore.Store
}

func (c *Reconciler) Reconcile(id interface{}) (bool, error) {
	return c.reconcile(id.(networktype.ID))
}

func (c *Reconciler) reconcile(id networktype.ID) (bool, error) {
	network, err := c.networks.Get(id)
	if err != nil {
		return false, err
	}

	// TODO: Changes should be removed when a network is removed

	// Ensure changes are created in the change store for each device in the network config.
	for _, networkChange := range network.Changes {
		change, err := c.changes.Get(network.GetChangeID(networkChange.DeviceID))
		if err != nil {
			return false, err
		}

		// If the change has not been created, create it.
		if change == nil {
			change = networkChange
			change.Status = changetype.Status_PENDING
			err = c.changes.Create(change)
			if err != nil {
				return false, err
			}
		}
	}

	// If the network status is APPLYING, ensure all device changes are in the APPLYING state.
	if network.Status == networktype.Status_APPLYING {
		status := networktype.Status_SUCCEEDED
		reason := networktype.Reason_ERROR
		for _, id := range network.GetChangeIDs() {
			change, err := c.changes.Get(id)
			if err != nil {
				return false, err
			}

			// If the change is PENDING then change it to APPLYING
			if change.Status == changetype.Status_PENDING {
				change.Status = changetype.Status_APPLYING
				status = networktype.Status_APPLYING
				err = c.changes.Update(change)
				if err != nil {
					return false, err
				}
			} else if change.Status == changetype.Status_APPLYING {
				// If the change is APPLYING then ensure the network status is APPLYING
				status = networktype.Status_APPLYING
			} else if change.Status == changetype.Status_FAILED {
				// If the change is FAILED then set the network to FAILED
				// If the change failure reason is UNAVAILABLE then all changes must be UNAVAILABLE, otherwise
				// the network must be failed with an ERROR.
				if status != networktype.Status_FAILED {
					switch change.Reason {
					case changetype.Reason_ERROR:
						reason = networktype.Reason_ERROR
					case changetype.Reason_UNAVAILABLE:
						reason = networktype.Reason_UNAVAILABLE
					}
				} else if reason == networktype.Reason_UNAVAILABLE && change.Reason == changetype.Reason_ERROR {
					reason = networktype.Reason_ERROR
				} else if reason == networktype.Reason_ERROR && change.Reason == changetype.Reason_UNAVAILABLE {
					reason = networktype.Reason_ERROR
				}
				status = networktype.Status_FAILED
			}
		}

		// If the status has changed, update the network config
		if network.Status != status || network.Reason != reason {
			network.Status = status
			network.Reason = reason
			if err := c.networks.Update(network); err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

var _ controller.Reconciler = &Reconciler{}
