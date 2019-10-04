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
	devicesnapshot "github.com/onosproject/onos-config/pkg/types/device/snapshot"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"strconv"
	"strings"
)

const separator = ":"

// ID is a network snapshot identifier type
type ID types.ID

// GetIndex returns the Index of the snapshot request
func (i ID) GetIndex() Index {
	indexStr := string(i)[strings.LastIndex(string(i), separator)+1:]
	index, _ := strconv.Atoi(indexStr)
	return Index(index)
}

// Index is a globally unique, monotonically increasing index of a network snapshot request
type Index uint64

// GetSnapshotID returns the network snapshot ID for the index
func (i Index) GetSnapshotID() ID {
	return ID(fmt.Sprintf("snapshot%s%d", separator, i))
}

// Revision is a network snapshot revision number
type Revision types.Revision

// GetDeviceSnapshotID gets the device snapshot ID for the given device/version
func (s *NetworkSnapshot) GetDeviceSnapshotID(deviceID device.ID, version string) devicesnapshot.ID {
	return devicesnapshot.NewID(types.ID(s.ID), deviceID, version)
}

// GetDeviceSnapshotIDs returns a list of snapshot IDs for the devices in the snapshot
func (s *NetworkSnapshot) GetDeviceSnapshotIDs() []devicesnapshot.ID {
	snapshotIDs := make([]devicesnapshot.ID, len(s.Devices))
	for i, device := range s.Devices {
		snapshotIDs[i] = s.GetDeviceSnapshotID(device.DeviceID, device.DeviceVersion)
	}
	return snapshotIDs
}
