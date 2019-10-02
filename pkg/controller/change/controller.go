package change

import (
	"github.com/onosproject/onos-config/pkg/controller"
	changestore "github.com/onosproject/onos-config/pkg/store/change"
	mastershipstore "github.com/onosproject/onos-config/pkg/store/mastership"
	changetype "github.com/onosproject/onos-config/pkg/types/change"
)

// NewController returns a new network controller
func NewController(mastership mastershipstore.Store, changes changestore.Store) *controller.Controller {
	c := controller.NewController()
	c.Filter(&controller.MastershipFilter{
		Store: mastership,
	})
	c.Watch(&Watcher{
		Store: changes,
	})
	c.Reconcile(&Reconciler{
		changes: changes,
	})
	return c
}

// Reconciler is the change reconciler
type Reconciler struct {
	changes changestore.Store
}

func (r *Reconciler) Reconcile(id interface{}) (bool, error) {
	return r.reconcile(id.(changetype.ID))
}

func (r *Reconciler) reconcile(id changetype.ID) (bool, error) {
	change, err := r.changes.Get(id)
	if err != nil {
		return false, err
	}

	// If the change is in the applying state, apply the change.
	if change.Status == changetype.Status_APPLYING {
		if err := r.applyChange(change); err != nil {
			return false, err
		}
	}
	return true, nil
}

// applyChange applies the given change to the device
func (r *Reconciler) applyChange(change *changetype.Change) error {
	change.Status = changetype.Status_FAILED
	change.Reason = changetype.Reason_UNAVAILABLE
	return r.changes.Update(change)
}

var _ controller.Reconciler = &Reconciler{}
