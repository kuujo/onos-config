package network

import (
	"github.com/onosproject/onos-config/pkg/controller"
	networkstore "github.com/onosproject/onos-config/pkg/store/network"
	networktype "github.com/onosproject/onos-config/pkg/types/network"
	"sync"
)

const queueSize = 100

// Watcher is a network watcher
type Watcher struct {
	Store networkstore.Store
	ch    chan *networktype.NetworkConfig
	mu    sync.Mutex
}

func (w *Watcher) Start(ch chan<- interface{}) error {
	defer close(ch)

	w.mu.Lock()
	if w.ch != nil {
		return nil
	}

	requestCh := make(chan *networktype.NetworkConfig, queueSize)
	w.ch = requestCh
	w.mu.Unlock()

	if err := w.Store.Watch(requestCh); err != nil {
		return err
	}

	for request := range requestCh {
		ch <- request.ID
	}
	return nil
}

func (w *Watcher) Stop() {
	w.mu.Lock()
	close(w.ch)
	w.mu.Unlock()
}

var _ controller.Watcher = &Watcher{}

// OwnerWatcher is a network owner watcher
type OwnerWatcher struct {
	Store networkstore.Store
	ch    chan *networktype.NetworkConfig
	mu    sync.Mutex
}

func (w *OwnerWatcher) Start(ch chan<- interface{}) error {
	defer close(ch)

	w.mu.Lock()
	if w.ch != nil {
		return nil
	}

	requestCh := make(chan *networktype.NetworkConfig, queueSize)
	w.ch = requestCh
	w.mu.Unlock()

	if err := w.Store.Watch(requestCh); err != nil {
		return err
	}

	for request := range requestCh {
		ch <- request.RequestID
	}
	return nil
}

func (w *OwnerWatcher) Stop() {
	w.mu.Lock()
	close(w.ch)
	w.mu.Unlock()
}

var _ controller.Watcher = &OwnerWatcher{}
