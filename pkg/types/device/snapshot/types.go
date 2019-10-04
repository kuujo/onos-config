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

package snapshot

import (
	"fmt"
	"github.com/onosproject/onos-config/pkg/types"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"strings"
)

const separator = ":"

// NewID returns a new device snapshot ID for the given device/version/index
func NewID(networkID types.ID, deviceID device.ID, version string) ID {
	return ID(fmt.Sprintf("%s%s%s%s%s", networkID, separator, deviceID, separator, version))
}

// ID is a snapshot identifier type
type ID types.ID

// GetDeviceID returns the ID of the device to which the snapshot applies
func (i ID) GetDeviceID() device.ID {
	return device.ID(string(i)[strings.Index(string(i), separator)+1 : strings.LastIndex(string(i), separator)])
}

// GetDeviceVersion returns the device version to which the snapshot applies
func (i ID) GetDeviceVersion() string {
	return string(i)[strings.LastIndex(string(i), separator)+1:]
}

// Index is the index of a snapshot
type Index uint64

// Revision is a network configuration revision number
type Revision types.Revision
