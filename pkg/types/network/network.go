package network

import (
	"fmt"
	"github.com/onosproject/onos-config/pkg/types/change"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
)

// ID is the identifier type for a network config
type ID string

// Revision is a network config revision number
type Revision uint64

// GetChangeIDs returns a list of change IDs for the network
func (c *NetworkConfig) GetChangeIDs() []change.ID {
	changes := c.GetChanges()
	if changes == nil {
		return nil
	}

	ids := make([]change.ID, len(changes))
	for i, change := range changes {
		ids[i] = c.GetChangeID(change.DeviceID)
	}
	return ids
}

// GetChangeID returns the ID for the change to the given device
func (c *NetworkConfig) GetChangeID(deviceID device.ID) change.ID {
	return change.ID(fmt.Sprintf("network-%d-%s", c.ID, deviceID))
}
