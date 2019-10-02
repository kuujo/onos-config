package controller

import (
	"github.com/onosproject/onos-config/pkg/store/cluster"
	leadershipstore "github.com/onosproject/onos-config/pkg/store/leadership"
	"sync"
)

// Activator is an interface for controlling the activation of a controller
type Activator interface {
	// Start starts the activator
	Start(ch chan<- bool) error

	// Stop stops the activator
	Stop()
}

// AlwaysActivator allows any node to handle a request
type AlwaysActivator struct {
}

func (a *AlwaysActivator) Start(ch chan<- bool) error {
	ch <- true
	return nil
}

func (a *AlwaysActivator) Stop() {

}

var _ Activator = &AlwaysActivator{}

// LeadershipActivator is an Activator for activating a controller on leadership
type LeadershipActivator struct {
	Store leadershipstore.Store
	ch    chan leadershipstore.Leadership
	mu    sync.Mutex
}

func (a *LeadershipActivator) Start(ch chan<- bool) error {
	a.mu.Lock()
	a.ch = make(chan leadershipstore.Leadership)
	a.mu.Unlock()

	if err := a.Store.Watch(a.ch); err != nil {
		return err
	}

	go func() {
		for leadership := range a.ch {
			if leadership.Leader == cluster.GetNodeID() {
				ch <- true
			} else {
				ch <- false
			}
		}
	}()
	return nil
}

func (a *LeadershipActivator) Stop() {
	a.mu.Lock()
	close(a.ch)
	a.mu.Unlock()
}

var _ Activator = &LeadershipActivator{}
