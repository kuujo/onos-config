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

func TestMastershipElection(t *testing.T) {
	node, conn := utils.StartLocalNode()

	store1, err := newLocalElection(topodevice.ID("test"), "a", conn)
	assert.NoError(t, err)

	store2, err := newLocalElection(topodevice.ID("test"), "b", conn)
	assert.NoError(t, err)

	store3, err := newLocalElection(topodevice.ID("test"), "c", conn)
	assert.NoError(t, err)

	mastership, err := store1.join()
	assert.NoError(t, err)
	assert.Equal(t, store1.NodeID(), mastership.Master)

	mastership, err = store2.join()
	assert.NoError(t, err)
	assert.Equal(t, store1.NodeID(), mastership.Master)

	mastership, err = store3.join()
	assert.NoError(t, err)
	assert.Equal(t, store1.NodeID(), mastership.Master)

	store2Ch := make(chan stream.Event)
	_, err = store2.watch(store2Ch)
	assert.NoError(t, err)

	store3Ch := make(chan stream.Event)
	_, err = store3.watch(store3Ch)
	assert.NoError(t, err)

	err = store1.Close()
	assert.NoError(t, err)

	event := <-store2Ch
	assert.Equal(t, cluster.NodeID("b"), event.Object.(Mastership).Master)

	master, err := store2.isMaster()
	assert.NoError(t, err)
	assert.True(t, master)

	event = <-store3Ch
	assert.Equal(t, cluster.NodeID("b"), event.Object.(Mastership).Master)

	master, err = store3.isMaster()
	assert.NoError(t, err)
	assert.False(t, master)

	err = store2.Close()
	assert.NoError(t, err)

	event = <-store3Ch
	assert.Equal(t, cluster.NodeID("c"), event.Object.(Mastership).Master)

	master, err = store3.isMaster()
	assert.NoError(t, err)
	assert.True(t, master)

	_ = store3.Close()
	_ = conn.Close()
	_ = node.Stop()
}
