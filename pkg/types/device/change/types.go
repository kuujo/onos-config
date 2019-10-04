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

package change

import (
	"fmt"
	"github.com/onosproject/onos-config/pkg/types"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"strconv"
	"strings"
)

// ID is a network configuration identifier type
type ID types.ID

// GetDeviceID returns the Device ID
func (i ID) GetDeviceID() device.ID {
	return device.ID(string(i)[strings.LastIndex(string(i), ":")+1:])
}

// GetDeviceVersion returns the device version
func (i ID) GetDeviceVersion() string {
	return string(i)[strings.Index(string(i), ":")+1 : strings.LastIndex(string(i), ":")]
}

// GetIndex returns the Index
func (i ID) GetIndex() Index {
	index, _ := strconv.Atoi(string(i)[strings.LastIndex(string(i), ":")+1:])
	return Index(index)
}

// Index is the index of a network configuration
type Index uint64

// GetID returns the ID for the index
func (i Index) GetID(deviceID device.ID, version string) ID {
	return ID(fmt.Sprintf("%s:%s:%d", deviceID, version, i))
}

// Revision is a network configuration revision number
type Revision types.Revision
