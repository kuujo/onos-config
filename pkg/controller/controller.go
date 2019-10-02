package controller

import (
	"container/list"
	"sync"
	"time"
)

// Watcher is implemented by controllers to implement watching for specific events
type Watcher interface {
	// Start starts watching for events
	Start(ch chan<- interface{}) error

	// Stop stops watching for events
	Stop()
}

// Reconciler reconciles objects
type Reconciler interface {
	// Reconcile is called to reconcile the state of an object
	Reconcile(interface{}) (bool, error)
}

// NewController creates a new controller
func NewController() *Controller {
	return &Controller{
		activator:  &AlwaysActivator{},
		watchers:   make([]Watcher, 0),
		retryQueue: list.New(),
	}
}

// Controller is the request controller
type Controller struct {
	mu         sync.Mutex
	activator  Activator
	filter     Filter
	watchers   []Watcher
	reconciler Reconciler
	retryQueue *list.List
}

// Activate sets an activator for the controller
func (c *Controller) Activate(activator Activator) {
	c.mu.Lock()
	c.activator = activator
	c.mu.Unlock()
}

// Filter sets a filter for the controller
func (c *Controller) Filter(filter Filter) {
	c.mu.Lock()
	c.filter = filter
	c.mu.Unlock()
}

// Watch adds a watcher to the controller
func (c *Controller) Watch(watcher Watcher) {
	c.mu.Lock()
	c.watchers = append(c.watchers, watcher)
	c.mu.Unlock()
}

// Reconcile sets the reconciler for the controller
func (c *Controller) Reconcile(reconciler Reconciler) {
	c.mu.Lock()
	c.reconciler = reconciler
	c.mu.Unlock()
}

// Start starts the request controller
func (c *Controller) Start() error {
	ch := make(chan bool)
	if err := c.activator.Start(ch); err != nil {
		return err
	}
	go func() {
		for activate := range ch {
			if activate {
				c.activate()
			} else {
				c.deactivate()
			}
		}
	}()
	return nil
}

// activate activates the controller
func (c *Controller) activate() {
	ch := make(chan interface{})
	for _, watcher := range c.watchers {
		go watcher.Start(ch)
	}
	go c.process(ch)
}

// deactivate deactivates the controller
func (c *Controller) deactivate() {
	for _, watcher := range c.watchers {
		watcher.Stop()
	}
}

// process processes the events from the given channel
func (c Controller) process(ch chan interface{}) {
	c.mu.Lock()
	filter := c.filter
	c.mu.Unlock()
	for id := range ch {
		if filter == nil || filter.Accept(id) {
			c.reconcile(id)
		}
	}
}

// retry retries the given list of requests
func (c *Controller) retry(requests *list.List) {
	element := requests.Front()
	for element != nil {
		c.reconcile(element.Value)
	}
}

// reconcile reconciles the given request ID until complete
func (c *Controller) reconcile(id interface{}) {
	c.mu.Lock()
	reconciler := c.reconciler
	c.mu.Unlock()

	iteration := 1
	for {
		c.mu.Lock()

		// First, check if any requests are pending in the retry queue
		if c.retryQueue.Len() > 0 {
			retries := c.retryQueue
			c.retryQueue = list.New()
			c.mu.Unlock()
			c.retry(retries)
			c.mu.Lock()
		}

		succeeded, err := reconciler.Reconcile(id)
		if err != nil {
			c.mu.Unlock()
			time.Sleep(time.Duration(iteration*2) * time.Millisecond)
		} else if !succeeded {
			c.retryQueue.PushBack(id)
			c.mu.Unlock()
			return
		} else {
			c.mu.Unlock()
		}
		iteration++
	}
}
