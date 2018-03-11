package flux

import (
	"sync"
)

type Dispatcher struct {
	m        sync.Mutex
	notifier NotifierInterface
	stores   []StoreInterface
}

// NewDispatcher creates a new dispatcher and registers the provided stores
func NewDispatcher(notifier NotifierInterface, stores ...StoreInterface) *Dispatcher {
	d := &Dispatcher{notifier: notifier}
	for _, s := range stores {
		d.Register(s)
	}
	return d
}

// Register adds a store to the dispatcher
func (d *Dispatcher) Register(store StoreInterface) {
	d.m.Lock()
	defer d.m.Unlock()
	d.stores = append(d.stores, store)
}

// Dispatch sends an action to all registered stores
func (d *Dispatcher) Dispatch(action ActionInterface) chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)
		payloads := make(map[StoreInterface]*Payload, len(d.stores))
		d.handle(action, payloads)
		d.notify(payloads)
	}()
	return done
}

func (d *Dispatcher) handle(action ActionInterface, payloads map[StoreInterface]*Payload) {
	d.m.Lock()
	defer d.m.Unlock()

	// Create a waitgroup and add the number of stores
	wg := sync.WaitGroup{}
	wg.Add(len(d.stores))

	// Create the loop detector
	loop := newLoopDetector()

	for _, store := range d.stores {
		payloads[store] = newPayload(action, store, payloads, loop, d.notifier)
	}

	for _, store := range d.stores {
		// The store will be used inside a goroutine
		store := store

		payload := payloads[store]

		go func() {
			// Start the store handler.
			finished := store.Handle(payload)

			// If we finished synchronously, we close the tracker's Done channel. If we are
			// still processing asynchronously, we leave it to be closed when the handler has
			// finished.
			if finished {
				close(payload.Done)
			}
		}()

		go func() {
			// We wait for the tracker to finish
			<-payload.finished()
			wg.Done()
		}()
	}
	wg.Wait()
}

func (d *Dispatcher) notify(payloads map[StoreInterface]*Payload) {
	// all stores have finished processing, so we fire a notification if needed
	var notify bool
	for _, p := range payloads {
		if p.notify {
			notify = true
			break
		}
	}
	if notify {
		<-d.notifier.Notify()
	}
}
