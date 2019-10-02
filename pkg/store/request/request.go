package request

import (
	"context"
	"fmt"
	"github.com/atomix/atomix-go-client/pkg/client/list"
	"github.com/atomix/atomix-go-client/pkg/client/primitive"
	"github.com/atomix/atomix-go-client/pkg/client/session"
	"github.com/atomix/atomix-go-local/pkg/atomix/local"
	"github.com/atomix/atomix-go-node/pkg/atomix"
	"github.com/atomix/atomix-go-node/pkg/atomix/registry"
	"github.com/gogo/protobuf/proto"
	"github.com/onosproject/onos-config/pkg/store/network"
	"github.com/onosproject/onos-config/pkg/store/utils"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"net"
	"time"
)

const requestTimeout = 15 * time.Second

// ID is a request identifier
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

	requests, err := group.GetList(context.Background(), "requests", session.WithTimeout(30*time.Second))
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		requests: requests,
		closer:   requests,
	}, nil
}

// NewLocalStore returns a new local device store
func NewLocalStore() (Store, error) {
	node, conn := startLocalNode()
	name := primitive.Name{
		Namespace: "local",
		Name:      "requests",
	}

	requests, err := list.New(context.Background(), name, []*grpc.ClientConn{conn})
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		requests: requests,
		closer:   &utils.NodeCloser{node},
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

	conn, err := grpc.DialContext(context.Background(), "requests", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		panic("Failed to dial devices")
	}
	return node, conn
}

// Store stores ConfigRequests
type Store interface {
	io.Closer

	// Get gets a configuration request
	Get(id ID) (*ConfigRequest, error)

	// ConfigRequest appends a configuration request
	Append(config *ConfigRequest) error

	// Replay replays configuration requests from the given ID
	Replay(ID, chan<- *ConfigRequest) error

	// Watch watches the store for changes
	Watch(chan<- *ConfigRequest) error
}

// atomixStore is the default implementation of the NetworkConfig store
type atomixStore struct {
	requests list.List
	closer   io.Closer
}

func (s *atomixStore) Get(id ID) (*ConfigRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	bytes, err := s.requests.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return decodeRequest(id, bytes)
}

func (s *atomixStore) Append(config *ConfigRequest) error {
	bytes, err := proto.Marshal(config)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	return s.requests.Append(ctx, bytes)
}

func (s *atomixStore) Replay(id ID, ch chan<- *ConfigRequest) error {
	itemsCh := make(chan []byte)
	go func() {
		i := int(id)
		for item := range itemsCh {
			request, err := decodeRequest(ID(i), item)
			if err == nil {
				ch <- request
			}
			i++
		}
	}()

	slice, err := s.requests.SliceFrom(context.Background(), int(id))
	if err != nil {
		return err
	}
	return slice.Items(context.Background(), itemsCh)
}

func (s *atomixStore) Watch(ch chan<- *ConfigRequest) error {
	watchCh := make(chan *list.Event)
	go func() {
		for event := range watchCh {
			request, err := decodeRequest(ID(event.Index), event.Value)
			if err == nil {
				ch <- request
			}
		}
	}()
	return s.requests.Watch(context.Background(), watchCh)
}

func (s *atomixStore) Close() error {
	_ = s.requests.Close()
	return s.closer.Close()
}

func decodeRequest(id ID, value []byte) (*ConfigRequest, error) {
	request := &ConfigRequest{}
	if err := proto.Unmarshal(value, request); err != nil {
		return nil, err
	}
	request.ID = id
	return request, nil
}
