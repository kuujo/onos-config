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
	"context"
	"errors"
	"github.com/atomix/atomix-go-client/pkg/client/counter"
	"github.com/atomix/atomix-go-client/pkg/client/map"
	"github.com/atomix/atomix-go-client/pkg/client/primitive"
	"github.com/atomix/atomix-go-client/pkg/client/session"
	"github.com/atomix/atomix-go-local/pkg/atomix/local"
	"github.com/atomix/atomix-go-node/pkg/atomix"
	"github.com/atomix/atomix-go-node/pkg/atomix/registry"
	"github.com/gogo/protobuf/proto"
	"github.com/onosproject/onos-config/pkg/store/utils"
	"github.com/onosproject/onos-config/pkg/types/network/change"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"net"
	"time"
)

const counterName = "config-ids"
const configsName = "configs"

// NewAtomixStore returns a new persistent Store
func NewAtomixStore() (Store, error) {
	client, err := utils.GetAtomixClient()
	if err != nil {
		return nil, err
	}

	group, err := client.GetGroup(context.Background(), utils.GetAtomixRaftGroup())
	if err != nil {
		return nil, err
	}

	ids, err := group.GetCounter(context.Background(), counterName, session.WithTimeout(30*time.Second))
	if err != nil {
		return nil, err
	}

	configs, err := group.GetMap(context.Background(), configsName, session.WithTimeout(30*time.Second))
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		ids:     ids,
		configs: configs,
		closer:  configs,
	}, nil
}

// NewLocalStore returns a new local device store
func NewLocalStore() (Store, error) {
	node, conn := startLocalNode()

	counterName := primitive.Name{
		Namespace: "local",
		Name:      counterName,
	}
	ids, err := counter.New(context.Background(), counterName, []*grpc.ClientConn{conn})
	if err != nil {
		return nil, err
	}

	configsName := primitive.Name{
		Namespace: "local",
		Name:      configsName,
	}
	configs, err := _map.New(context.Background(), configsName, []*grpc.ClientConn{conn})
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		ids:     ids,
		configs: configs,
		closer:  utils.NewNodeCloser(node),
	}, nil
}

// startLocalNode starts a single local node
func startLocalNode() (*atomix.Node, *grpc.ClientConn) {
	lis := bufconn.Listen(1024 * 1024)
	node := local.NewNode(lis, registry.Registry)
	_ = node.Start()

	dialer := func(ctx context.Context, address string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), configsName, grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		panic("Failed to dial network configurations")
	}
	return node, conn
}

// Store stores NetworkConfig changes
type Store interface {
	io.Closer

	// Get gets a network configuration
	Get(id change.ID) (*change.NetworkChange, error)

	// Create creates a new network configuration
	Create(config *change.NetworkChange) error

	// Update updates an existing network configuration
	Update(config *change.NetworkChange) error

	// Delete deletes a network configuration
	Delete(config *change.NetworkChange) error

	// List lists network configurations
	List(chan<- *change.NetworkChange) error

	// Watch watches the network configuration store for changes
	Watch(chan<- *change.NetworkChange) error
}

// atomixStore is the default implementation of the NetworkConfig store
type atomixStore struct {
	ids     counter.Counter
	configs _map.Map
	closer  io.Closer
}

func (s *atomixStore) Get(id change.ID) (*change.NetworkChange, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	entry, err := s.configs.Get(ctx, string(id))
	if err != nil {
		return nil, err
	} else if entry == nil {
		return nil, nil
	}
	return decodeConfig(entry)
}

func (s *atomixStore) Create(config *change.NetworkChange) error {
	if config.Revision != 0 {
		return errors.New("not a new object")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	id, err := s.ids.Increment(ctx, 1)
	cancel()
	if err != nil {
		return err
	}

	config.Index = change.Index(id)

	bytes, err := proto.Marshal(config)
	if err != nil {
		return err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	entry, err := s.configs.Put(ctx, string(config.ID), bytes, _map.IfNotSet())
	cancel()
	if err != nil {
		return err
	}

	config.Revision = change.Revision(entry.Version)
	config.Created = entry.Created
	config.Updated = entry.Updated
	return nil
}

func (s *atomixStore) Update(config *change.NetworkChange) error {
	if config.Revision == 0 {
		return errors.New("not a stored object")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bytes, err := proto.Marshal(config)
	if err != nil {
		return err
	}

	entry, err := s.configs.Put(ctx, string(config.ID), bytes, _map.IfVersion(int64(config.Revision)))
	if err != nil {
		return err
	}

	config.Revision = change.Revision(entry.Version)
	config.Updated = entry.Updated
	return nil
}

func (s *atomixStore) Delete(config *change.NetworkChange) error {
	if config.Revision == 0 {
		return errors.New("not a stored object")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	entry, err := s.configs.Remove(ctx, string(config.ID), _map.IfVersion(int64(config.Revision)))
	if err != nil {
		return err
	}
	config.Updated = entry.Updated
	return nil
}

func (s *atomixStore) List(ch chan<- *change.NetworkChange) error {
	mapCh := make(chan *_map.Entry)
	if err := s.configs.Entries(context.Background(), mapCh); err != nil {
		return err
	}

	go func() {
		defer close(ch)
		for entry := range mapCh {
			if device, err := decodeConfig(entry); err == nil {
				ch <- device
			}
		}
	}()
	return nil
}

func (s *atomixStore) Watch(ch chan<- *change.NetworkChange) error {
	mapCh := make(chan *_map.Event)
	if err := s.configs.Watch(context.Background(), mapCh, _map.WithReplay()); err != nil {
		return err
	}

	go func() {
		defer close(ch)
		for event := range mapCh {
			if config, err := decodeConfig(event.Entry); err == nil {
				ch <- config
			}
		}
	}()
	return nil
}

func (s *atomixStore) Close() error {
	_ = s.configs.Close()
	return s.closer.Close()
}

func decodeConfig(entry *_map.Entry) (*change.NetworkChange, error) {
	conf := &change.NetworkChange{}
	if err := proto.Unmarshal(entry.Value, conf); err != nil {
		return nil, err
	}
	conf.ID = change.ID(entry.Key)
	conf.Revision = change.Revision(entry.Version)
	conf.Created = entry.Created
	conf.Updated = entry.Updated
	return conf, nil
}