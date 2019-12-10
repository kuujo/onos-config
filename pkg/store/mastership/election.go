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
	"context"
	"fmt"
	"github.com/atomix/atomix-go-client/pkg/client"
	"github.com/atomix/atomix-go-client/pkg/client/election"
	"github.com/atomix/atomix-go-client/pkg/client/primitive"
	"github.com/atomix/atomix-go-client/pkg/client/session"
	"github.com/onosproject/onos-config/pkg/store/cluster"
	"github.com/onosproject/onos-config/pkg/store/stream"
	topodevice "github.com/onosproject/onos-topo/api/device"
	"google.golang.org/grpc"
	"io"
	"time"
)

// newAtomixElection returns a new persistent device mastership election
func newAtomixElection(deviceID topodevice.ID, group *client.PartitionGroup) (deviceMastershipElection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	election, err := group.GetElection(ctx, fmt.Sprintf("mastership-%s", deviceID), session.WithID(string(cluster.GetNodeID())), session.WithTimeout(15*time.Second))
	cancel()
	if err != nil {
		return nil, err
	}
	return newDeviceMastershipElection(deviceID, election)
}

// newLocalElection returns a new local device mastership election
func newLocalElection(deviceID topodevice.ID, nodeID cluster.NodeID, conn *grpc.ClientConn) (deviceMastershipElection, error) {
	name := primitive.Name{
		Namespace: "local",
		Name:      fmt.Sprintf("mastership-%s", deviceID),
	}
	election, err := election.New(context.Background(), name, []*grpc.ClientConn{conn}, session.WithID(string(nodeID)), session.WithTimeout(15*time.Second))
	if err != nil {
		return nil, err
	}
	return newDeviceMastershipElection(deviceID, election)
}

// newDeviceMastershipElection creates and enters a new device mastership election
func newDeviceMastershipElection(deviceID topodevice.ID, election election.Election) (deviceMastershipElection, error) {
	return &atomixDeviceMastershipElection{
		deviceID: deviceID,
		election: election,
		watchers: make([]chan<- Mastership, 0, 1),
	}, nil
}

// deviceMastershipElection is an election for a single device mastership
type deviceMastershipElection interface {
	io.Closer

	// NodeID returns the local node identifier used in the election
	NodeID() cluster.NodeID

	// DeviceID returns the device for which this election provides mastership
	DeviceID() topodevice.ID

	// join joins the mastership election
	join() (Mastership, error)

	// leave leaves thee mastership election
	leave() (Mastership, error)

	// isMaster returns a bool indicating whether the local node is the master for the device
	isMaster() (bool, error)

	// watch watches the election for changes
	watch(ch chan<- stream.Event) (stream.Context, error)
}

// atomixDeviceMastershipElection is a persistent device mastership election
type atomixDeviceMastershipElection struct {
	deviceID topodevice.ID
	election election.Election
	watchers []chan<- Mastership
}

func (e *atomixDeviceMastershipElection) NodeID() cluster.NodeID {
	return cluster.NodeID(e.election.ID())
}

func (e *atomixDeviceMastershipElection) DeviceID() topodevice.ID {
	return e.deviceID
}

func (e *atomixDeviceMastershipElection) join() (Mastership, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	term, err := e.election.Enter(ctx)
	if err != nil {
		return Mastership{}, err
	}
	return Mastership{
		Device: e.deviceID,
		Term:   Term(term.ID),
		Master: cluster.NodeID(term.Leader),
	}, nil
}

func (e *atomixDeviceMastershipElection) leave() (Mastership, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	term, err := e.election.Leave(ctx)
	if err != nil {
		return Mastership{}, err
	}
	return Mastership{
		Device: e.deviceID,
		Term:   Term(term.ID),
		Master: cluster.NodeID(term.Leader),
	}, nil
}

func (e *atomixDeviceMastershipElection) isMaster() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	term, err := e.election.GetTerm(ctx)
	if err != nil {
		return false, err
	}
	return term.Leader == e.election.ID(), nil
}

func (e *atomixDeviceMastershipElection) watch(ch chan<- stream.Event) (stream.Context, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	electionCh := make(chan *election.Event)
	if err := e.election.Watch(ctx, electionCh); err != nil {
		return nil, err
	}

	go func() {
		defer close(ch)
		for event := range electionCh {
			ch <- stream.Event{
				Type: stream.Updated,
				Object: Mastership{
					Device: e.deviceID,
					Term:   Term(event.Term.ID),
					Master: cluster.NodeID(event.Term.Leader),
				},
			}
		}
	}()
	return stream.NewCancelContext(cancel), nil
}

func (e *atomixDeviceMastershipElection) Close() error {
	return e.election.Close()
}

var _ deviceMastershipElection = &atomixDeviceMastershipElection{}
