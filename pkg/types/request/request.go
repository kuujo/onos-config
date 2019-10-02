package request

import (
	"fmt"
	"github.com/onosproject/onos-config/pkg/types/network"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
)

// ID is the identifier type for a request
type ID uint64

// GetNetworkID gets the network ID for the request
func (r *ConfigRequest) GetNetworkID() network.ID {
	return network.ID(fmt.Sprintf("%d", r.ID))
}

// GetDevices gets a list of devices changed by the configuration
func (r *ConfigRequest) GetDevices() []device.ID {
	devices := make([]device.ID, len(r.Changes))
	for i := 0; i < len(r.Changes); i++ {
		devices[i] = r.Changes[i].DeviceID
	}
	return devices
}
