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

// Package change defines change records for tracking device configuration changes.
package change

import (
	"context"
	"errors"
	"github.com/atomix/atomix-go-client/pkg/client/map"
	"github.com/atomix/atomix-go-client/pkg/client/primitive"
	"github.com/atomix/atomix-go-client/pkg/client/session"
	"github.com/atomix/atomix-go-local/pkg/atomix/local"
	"github.com/atomix/atomix-go-node/pkg/atomix"
	"github.com/atomix/atomix-go-node/pkg/atomix/registry"
	"github.com/gogo/protobuf/proto"
	"github.com/onosproject/onos-config/pkg/store/utils"
	"github.com/onosproject/onos-config/pkg/types/change"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"net"
	"time"
)

const primitiveName = "change"

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

	configs, err := group.GetMap(context.Background(), primitiveName, session.WithTimeout(30*time.Second))
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		configs: configs,
		closer:  configs,
	}, nil
}

// NewLocalStore returns a new local device store
func NewLocalStore() (Store, error) {
	node, conn := startLocalNode()
	name := primitive.Name{
		Namespace: "local",
		Name:      primitiveName,
	}

	configs, err := _map.New(context.Background(), name, []*grpc.ClientConn{conn})
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		configs: configs,
		closer:  &utils.NodeCloser{node},
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

	conn, err := grpc.DialContext(context.Background(), primitiveName, grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		panic("Failed to dial network configurations")
	}
	return node, conn
}

// Store stores DeviceChanges
type Store interface {
	io.Closer

	// Get gets a device change
	Get(id change.ID) (*change.Change, error)

	// Create creates a new device change
	Create(config *change.Change) error

	// Update updates an existing device change
	Update(config *change.Change) error

	// Delete deletes a device change
	Delete(config *change.Change) error

	// List lists device change
	List(chan<- *change.Change) error

	// Watch watches the device change store for changes
	Watch(chan<- *change.Change) error
}

// atomixStore is the default implementation of the NetworkConfig store
type atomixStore struct {
	configs _map.Map
	closer  io.Closer
}

func (s *atomixStore) Get(id change.ID) (*change.Change, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	kv, err := s.configs.Get(ctx, string(id))
	if err != nil {
		return nil, err
	} else if kv == nil {
		return nil, nil
	}
	return decodeChange(kv.Key, kv.Value, kv.Version)
}

func (s *atomixStore) Create(config *change.Change) error {
	if config.Revision != 0 {
		return errors.New("not a new object")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	config.Created = time.Now()
	config.Updated = time.Now()
	bytes, err := proto.Marshal(config)
	if err != nil {
		return err
	}

	kv, err := s.configs.Put(ctx, string(config.ID), bytes, _map.IfVersion(0))
	if err != nil {
		return err
	}

	config.Revision = change.Revision(kv.Version)
	return nil
}

func (s *atomixStore) Update(config *change.Change) error {
	if config.Revision == 0 {
		return errors.New("not a stored object")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	config.Updated = time.Now()
	bytes, err := proto.Marshal(config)
	if err != nil {
		return err
	}

	kv, err := s.configs.Put(ctx, string(config.ID), bytes, _map.IfVersion(int64(config.Revision)))
	if err != nil {
		return err
	}

	config.Revision = change.Revision(kv.Version)
	return nil
}

func (s *atomixStore) Delete(config *change.Change) error {
	if config.Revision == 0 {
		return errors.New("not a stored object")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := s.configs.Remove(ctx, string(config.ID), _map.IfVersion(int64(config.Revision)))
	return err
}

func (s *atomixStore) List(ch chan<- *change.Change) error {
	mapCh := make(chan *_map.KeyValue)
	if err := s.configs.Entries(context.Background(), mapCh); err != nil {
		return err
	}

	go func() {
		defer close(ch)
		for kv := range mapCh {
			if device, err := decodeChange(kv.Key, kv.Value, kv.Version); err == nil {
				ch <- device
			}
		}
	}()
	return nil
}

func (s *atomixStore) Watch(ch chan<- *change.Change) error {
	mapCh := make(chan *_map.Event)
	if err := s.configs.Watch(context.Background(), mapCh, _map.WithReplay()); err != nil {
		return err
	}

	go func() {
		defer close(ch)
		for event := range mapCh {
			if config, err := decodeChange(event.Key, event.Value, event.Version); err == nil {
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

func decodeChange(key string, value []byte, version int64) (*change.Change, error) {
	config := &change.Change{}
	if err := proto.Unmarshal(value, config); err != nil {
		return nil, err
	}
	config.ID = change.ID(key)
	config.Revision = change.Revision(version)
	return config, nil
}
