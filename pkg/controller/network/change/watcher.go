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
	"github.com/onosproject/onos-config/pkg/controller"
	changestore "github.com/onosproject/onos-config/pkg/store/network/change"
	"github.com/onosproject/onos-config/pkg/types"
	changetype "github.com/onosproject/onos-config/pkg/types/network/change"
	"sync"
)

const queueSize = 100

// Watcher is a network watcher
type Watcher struct {
	Store changestore.Store
	ch    chan *changetype.NetworkChange
	mu    sync.Mutex
}

// Start starts the config watcher
func (w *Watcher) Start(ch chan<- types.ID) error {
	defer close(ch)

	w.mu.Lock()
	if w.ch != nil {
		return nil
	}

	configCh := make(chan *changetype.NetworkChange, queueSize)
	w.ch = configCh
	w.mu.Unlock()

	if err := w.Store.Watch(configCh); err != nil {
		return err
	}

	for request := range configCh {
		ch <- types.ID(request.ID)
	}
	return nil
}

// Stop stops the config watcher
func (w *Watcher) Stop() {
	w.mu.Lock()
	close(w.ch)
	w.mu.Unlock()
}

var _ controller.Watcher = &Watcher{}