package mastership

import (
	"context"
	"fmt"
	"github.com/atomix/atomix-go-client/pkg/client"
	"github.com/atomix/atomix-go-client/pkg/client/election"
	"github.com/atomix/atomix-go-client/pkg/client/primitive"
	"github.com/atomix/atomix-go-local/pkg/atomix/local"
	"github.com/atomix/atomix-go-node/pkg/atomix"
	"github.com/atomix/atomix-go-node/pkg/atomix/registry"
	"github.com/onosproject/onos-config/pkg/store/cluster"
	"github.com/onosproject/onos-config/pkg/store/utils"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"net"
	"sync"
	"time"
)

// Term is a monotonically increasing mastership term
type Term uint64

// Store is the device mastership store
type Store interface {
	// IsMaster returns a boolean indicating whether the local node is the master for the given device
	IsMaster(id device.ID) (bool, error)

	// Watch watches the store for mastership changes
	Watch(device.ID, chan<- Mastership) error
}

// Mastership contains information about a device mastership term
type Mastership struct {
	Device device.ID
	Term   Term
	Master cluster.NodeID
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

	var closer io.Closer
	return &atomixStore{
		electionFactory: func(id device.ID) (election.Election, error) {
			return group.GetElection(context.Background(), fmt.Sprintf("mastership-%s", id))
		},
		elections:   make(map[device.ID]election.Election),
		masterships: make(map[device.ID]Mastership),
		closer:      closer,
	}, nil
}

// NewLocalStore returns a new local device store
func NewLocalStore() (Store, error) {
	node, conn := startLocalNode()
	return &atomixStore{
		electionFactory: func(id device.ID) (election.Election, error) {
			name := primitive.Name{
				Namespace: "local",
				Name:      fmt.Sprintf("mastership-%s", id),
			}
			return election.New(context.Background(), name, []*grpc.ClientConn{conn})
		},
		elections:   make(map[device.ID]election.Election),
		masterships: make(map[device.ID]Mastership),
		closer:      &utils.NodeCloser{node},
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

	conn, err := grpc.DialContext(context.Background(), "mastership", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		panic("Failed to dial network configurations")
	}
	return node, conn
}

// atomixStore is the default implementation of the NetworkConfig store
type atomixStore struct {
	electionFactory func(device.ID) (election.Election, error)
	group           *client.PartitionGroup
	closer          io.Closer
	elections       map[device.ID]election.Election
	masterships     map[device.ID]Mastership
	mu              sync.RWMutex
}

func (s *atomixStore) IsMaster(deviceID device.ID) (bool, error) {
	s.mu.RLock()
	mastership, ok := s.masterships[deviceID]
	if ok {
		defer s.mu.RUnlock()
		return mastership.Master == cluster.GetNodeID(), nil
	}
	s.mu.RUnlock()

	// Create a new election from the factory
	election, err := s.electionFactory(deviceID)
	if err != nil {
		return false, err
	}

	// Store the election session and enter the local node into the election
	s.mu.Lock()
	if _, ok := s.elections[deviceID]; !ok {
		s.elections[deviceID] = election
		s.mu.Unlock()

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		term, err := election.Enter(ctx)
		if err != nil {
			return false, err
		}
		return s.setTerm(deviceID, term).Master == cluster.GetNodeID(), nil
	}

	election.Close()
	election = s.elections[deviceID]
	s.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Get the current term for the election
	term, err := election.GetTerm(ctx)
	if err != nil {
		return false, err
	} else if term != nil {
		return s.setTerm(deviceID, term).Master == cluster.GetNodeID(), nil
	}
	return false, nil
}

func (s *atomixStore) setTerm(deviceID device.ID, term *election.Term) Mastership {
	s.mu.Lock()
	defer s.mu.Unlock()

	var mastershipTerm Term
	var master cluster.NodeID
	if term != nil {
		mastershipTerm = Term(term.ID)
		master = cluster.NodeID(term.Leader)
	}
	mastership := Mastership{
		Device: deviceID,
		Term:   mastershipTerm,
		Master: master,
	}
	s.masterships[deviceID] = mastership
	return mastership
}

func (s *atomixStore) Watch(deviceID device.ID, ch chan<- Mastership) error {
	s.mu.RLock()
	mastershipElection, ok := s.elections[deviceID]
	s.mu.RUnlock()
	if !ok {
		mastershipElection, err := s.electionFactory(deviceID)
		if err != nil {
			return err
		}

		// Store the election session and ensure the session was not duplicated
		s.mu.Lock()
		if _, ok := s.elections[deviceID]; !ok {
			s.elections[deviceID] = mastershipElection
		} else {
			_ = mastershipElection.Close()
			mastershipElection = s.elections[deviceID]
		}
		s.mu.Unlock()
	} else {
		s.mu.RUnlock()
	}

	electionCh := make(chan *election.Event)
	go func() {
		for event := range electionCh {
			ch <- s.setTerm(deviceID, &event.Term)
		}
	}()
	return mastershipElection.Watch(context.Background(), electionCh)
}

func (s *atomixStore) Close() error {
	for _, election := range s.elections {
		_ = election.Close()
	}
	return s.closer.Close()
}
