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

package mastership

import (
	"github.com/onosproject/onos-config/pkg/store/cluster"
	"github.com/onosproject/onos-config/pkg/store/stream"
	"github.com/onosproject/onos-config/pkg/store/utils"
	topodevice "github.com/onosproject/onos-topo/api/device"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMastershipStore(t *testing.T) {
	node, conn := utils.StartLocalNode()

	node1 := cluster.NodeID("node1")
	node2 := cluster.NodeID("node2")
	node3 := cluster.NodeID("node3")

	device1 := topodevice.ID("device1")
	device2 := topodevice.ID("device2")

	// Create three stores for three different nodes
	store1, err := newLocalStore(node1, conn)
	assert.NoError(t, err)

	store2, err := newLocalStore(node2, conn)
	assert.NoError(t, err)

	store3, err := newLocalStore(node3, conn)
	assert.NoError(t, err)

	// Verify that the first node that checks mastership for a device wins the election
	// and no other node believes itself to be the master
	mastership, err := store1.Join(device1)
	assert.NoError(t, err)
	assert.Equal(t, store1.NodeID(), mastership.Master)

	mastership, err = store2.Join(device1)
	assert.NoError(t, err)
	assert.Equal(t, store1.NodeID(), mastership.Master)

	mastership, err = store3.Join(device1)
	assert.NoError(t, err)
	assert.Equal(t, store1.NodeID(), mastership.Master)

	// Verify that listening for events for a device enters a node into the device mastership election
	store2Ch2 := make(chan stream.Event)
	_, err = store2.Watch(device2, store2Ch2)
	assert.NoError(t, err)

	// Verify that the watching node is the master
	mastership, err = store2.Join(device2)
	assert.NoError(t, err)
	assert.Equal(t, store2.NodeID(), mastership.Master)

	mastership, err = store1.Join(device2)
	assert.NoError(t, err)
	assert.Equal(t, store2.NodeID(), mastership.Master)

	mastership, err = store3.Join(device2)
	assert.NoError(t, err)
	assert.Equal(t, store2.NodeID(), mastership.Master)

	master, err := store3.IsMaster(device1)
	assert.NoError(t, err)
	assert.False(t, master)

	// Listen for device1 events on remaining nodes
	store2Ch1 := make(chan stream.Event)
	_, err = store2.Watch(device1, store2Ch1)
	assert.NoError(t, err)
	store3Ch1 := make(chan stream.Event)
	_, err = store3.Watch(device1, store3Ch1)
	assert.NoError(t, err)

	// Close node 1 (the master for device 1) and verify a mastership change occurs on nodes 2 and 3
	err = store1.Close()
	assert.NoError(t, err)

	// node2 should now be the master for device1
	event := <-store2Ch1
	assert.Equal(t, device1, event.Object.(Mastership).Device)
	assert.Equal(t, node2, event.Object.(Mastership).Master)

	master, err = store2.IsMaster(device1)
	assert.NoError(t, err)
	assert.True(t, master)

	// node3 should not be the master for device1
	event = <-store3Ch1
	assert.Equal(t, device1, event.Object.(Mastership).Device)
	assert.Equal(t, node2, event.Object.(Mastership).Master)

	master, err = store3.IsMaster(device1)
	assert.NoError(t, err)
	assert.False(t, master)

	// Close node2 and verify mastership for both devices changes
	err = store2.Close()
	assert.NoError(t, err)

	// node3 should now be the master for device1
	event = <-store3Ch1
	assert.Equal(t, device1, event.Object.(Mastership).Device)
	assert.Equal(t, node3, event.Object.(Mastership).Master)

	master, err = store3.IsMaster(device1)
	assert.NoError(t, err)
	assert.True(t, master)

	master, err = store3.IsMaster(device2)
	assert.NoError(t, err)
	assert.True(t, master)

	_ = store3.Close()
	_ = conn.Close()
	_ = node.Stop()
}
