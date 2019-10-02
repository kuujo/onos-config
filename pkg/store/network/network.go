package network

import (
	"context"
	"errors"
	"fmt"
	"github.com/atomix/atomix-go-client/pkg/client/map"
	"github.com/atomix/atomix-go-client/pkg/client/primitive"
	"github.com/atomix/atomix-go-client/pkg/client/session"
	"github.com/atomix/atomix-go-local/pkg/atomix/local"
	"github.com/atomix/atomix-go-node/pkg/atomix"
	"github.com/atomix/atomix-go-node/pkg/atomix/registry"
	"github.com/gogo/protobuf/proto"
	"github.com/onosproject/onos-config/pkg/store/change"
	"github.com/onosproject/onos-config/pkg/store/utils"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"net"
	"time"
)

const primitiveName = "network"

// ID is a network configuration identifier
type ID string

// Revision is the revision number
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

// Store stores NetworkConfig changes
type Store interface {
	io.Closer

	// Get gets a network configuration
	Get(id ID) (*NetworkConfig, error)

	// Create creates a new network configuration
	Create(config *NetworkConfig) error

	// Update updates an existing network configuration
	Update(config *NetworkConfig) error

	// Delete deletes a network configuration
	Delete(config *NetworkConfig) error

	// List lists network configurations
	List(chan<- *NetworkConfig) error

	// Watch watches the network configuration store for changes
	Watch(chan<- *NetworkConfig) error
}

// atomixStore is the default implementation of the NetworkConfig store
type atomixStore struct {
	configs _map.Map
	closer  io.Closer
}

func (s *atomixStore) Get(id ID) (*NetworkConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	kv, err := s.configs.Get(ctx, string(id))
	if err != nil {
		return nil, err
	} else if kv == nil {
		return nil, nil
	}
	return decodeConfig(kv.Key, kv.Value, kv.Version)
}

func (s *atomixStore) Create(config *NetworkConfig) error {
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

	config.Revision = Revision(kv.Version)
	return nil
}

func (s *atomixStore) Update(config *NetworkConfig) error {
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

	config.Revision = Revision(kv.Version)
	return nil
}

func (s *atomixStore) Delete(config *NetworkConfig) error {
	if config.Revision == 0 {
		return errors.New("not a stored object")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	config.Updated = time.Now()
	_, err := s.configs.Remove(ctx, string(config.ID), _map.IfVersion(int64(config.Revision)))
	return err
}

func (s *atomixStore) List(ch chan<- *NetworkConfig) error {
	mapCh := make(chan *_map.KeyValue)
	if err := s.configs.Entries(context.Background(), mapCh); err != nil {
		return err
	}

	go func() {
		defer close(ch)
		for kv := range mapCh {
			if device, err := decodeConfig(kv.Key, kv.Value, kv.Version); err == nil {
				ch <- device
			}
		}
	}()
	return nil
}

func (s *atomixStore) Watch(ch chan<- *NetworkConfig) error {
	mapCh := make(chan *_map.Event)
	if err := s.configs.Watch(context.Background(), mapCh, _map.WithReplay()); err != nil {
		return err
	}

	go func() {
		defer close(ch)
		for event := range mapCh {
			if config, err := decodeConfig(event.Key, event.Value, event.Version); err == nil {
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

func decodeConfig(key string, value []byte, version int64) (*NetworkConfig, error) {
	config := &NetworkConfig{}
	if err := proto.Unmarshal(value, config); err != nil {
		return nil, err
	}
	config.ID = ID(key)
	config.Revision = Revision(version)
	return config, nil
}
