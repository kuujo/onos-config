package device

import (
	"context"
	"github.com/atomix/atomix-go-client/pkg/client/util"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"google.golang.org/grpc"
	"io"
	"time"
)

const topoAddress = "onos-topo:5150"

// Store is a device store
type Store interface {
	// Get gets a device by ID
	Get(device.ID) (*device.Device, error)

	// Watch watches the device store for changes
	Watch(chan<- *device.Device) error
}

// NewTopoStore returns a new topo-based device store
func NewTopoStore(opts ...grpc.DialOption) (Store, error) {
	opts = append(opts,
		grpc.WithUnaryInterceptor(util.RetryingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(util.RetryingStreamClientInterceptor(100*time.Millisecond)))
	conn, err := getTopoConn(opts...)
	if err != nil {
		return nil, err
	}

	client := device.NewDeviceServiceClient(conn)
	return &topoStore{
		client: client,
	}, nil
}

// A device Store that uses the topo service to propagate devices
type topoStore struct {
	client device.DeviceServiceClient
}

func (s *topoStore) Get(id device.ID) (*device.Device, error) {
	response, err := s.client.Get(&device.GetRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return response.Device, nil
}

func (s *topoStore) Watch(ch chan<- *device.Device) error {
	list, err := s.client.List(context.Background(), &device.ListRequest{
		Subscribe: true,
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			response, err := list.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			ch <- response.Device
		}
	}()
	return nil
}

// getTopoConn gets a gRPC connection to the topology service
func getTopoConn(opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(topoAddress, opts...)
}
