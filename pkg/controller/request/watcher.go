package request

import (
	"github.com/onosproject/onos-config/pkg/controller"
	requeststore "github.com/onosproject/onos-config/pkg/store/request"
	requesttype "github.com/onosproject/onos-config/pkg/types/request"
	"sync"
)

const queueSize = 100

// Watcher is a request watcher
type Watcher struct {
	Store requeststore.Store
	ch    chan *requesttype.ConfigRequest
	mu    sync.Mutex
}

func (w *Watcher) Start(ch chan<- interface{}) error {
	defer close(ch)

	w.mu.Lock()
	if w.ch != nil {
		return nil
	}

	requestCh := make(chan *requesttype.ConfigRequest, queueSize)
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
