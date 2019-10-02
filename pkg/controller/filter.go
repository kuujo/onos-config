package controller

import (
	mastershipstore "github.com/onosproject/onos-config/pkg/store/mastership"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
)

// Filter filters individual events for a node
type Filter interface {
	// Accept indicates whether to accept the given object
	Accept(id interface{}) bool
}

// MastershipFilter activates a controller on acquiring mastership
type MastershipFilter struct {
	mastership mastershipstore.Store
}

// Accept accepts the given ID if the local node is the master
func (f *MastershipFilter) Accept(id interface{}) bool {
	master, err := f.mastership.IsMaster(id.(device.ID))
	if err != nil {
		return false
	}
	return master
}

var _ Filter = &MastershipFilter{}
