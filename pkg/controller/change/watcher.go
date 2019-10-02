package change

import (
	"github.com/onosproject/onos-config/pkg/controller"
	changestore "github.com/onosproject/onos-config/pkg/store/change"
	"sync"
)

const queueSize = 100

// Watcher is a change watcher
type Watcher struct {
	Store changestore.Store
	ch    chan *changestore.Change
	mu    sync.Mutex
}

func (w *Watcher) Start(ch chan<- interface{}) error {
	defer close(ch)

	w.mu.Lock()
	if w.ch != nil {
		return nil
	}

	requestCh := make(chan *changestore.Change, queueSize)
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

// OwnerWatcher is a change network watcher
type OwnerWatcher struct {
	Store changestore.Store
	ch    chan *changestore.Change
	mu    sync.Mutex
}

func (w *OwnerWatcher) Start(ch chan<- interface{}) error {
	defer close(ch)

	w.mu.Lock()
	if w.ch != nil {
		return nil
	}

	requestCh := make(chan *changestore.Change, queueSize)
	w.ch = requestCh
	w.mu.Unlock()

	if err := w.Store.Watch(requestCh); err != nil {
		return err
	}

	for request := range requestCh {
		ch <- request.NetworkID
	}
	return nil
}

func (w *OwnerWatcher) Stop() {
	w.mu.Lock()
	close(w.ch)
	w.mu.Unlock()
}

var _ controller.Watcher = &OwnerWatcher{}
