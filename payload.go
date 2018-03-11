package flux

import (
	"fmt"
	"sync"
)

type Payload struct {
	Action ActionInterface
	Done   chan struct{}

	loop        *loopDetector
	store       StoreInterface
	environment map[StoreInterface]*Payload
	complete    bool
	m           sync.Mutex
	notify      bool
	notifier    NotifierInterface

	// subscribers is a list of channels that will be closed when this payload has finished processing.
	// Each channel corresponds to another store that is waiting for this store to finish.
	subscribers []chan struct{}
}

func newPayload(action ActionInterface, store StoreInterface, environment map[StoreInterface]*Payload, loop *loopDetector, notifier NotifierInterface) *Payload {
	p := &Payload{
		Action:      action,
		Done:        make(chan struct{}),
		store:       store,
		environment: environment,
		loop:        loop,
		notifier:    notifier,
	}
	go p.monitor()
	return p
}

func (p *Payload) Notify() {
	p.notify = true
}

// monitor waits for the Done channel to be closed and notifies the subscribers that the payload
// has finished processing
func (p *Payload) monitor() {

	// Wait for the Done channel to be closed
	if _, open := <-p.Done; open {
		panic("never send on the Done channel - only close it")
	}

	// We don't want finished() to run at the same time as this function, so we use a mutex
	p.m.Lock()
	defer p.m.Unlock()

	// We mark the tracker as finished, and close the signal channel for all subscribers.
	p.complete = true
	for _, s := range p.subscribers {
		close(s)
	}
}

// Wait is used by the store handler to wait for other stores to finish.
func (p *Payload) Wait(stores ...StoreInterface) {

	// Check to see if we have a wait loop.
	if found, loopStore := p.loop.request(p.store, stores...); found {
		panic(fmt.Errorf("%T and %T are waiting for each other", p.store, loopStore))
	}

	// Inform the loop detector that we've finished waiting.
	defer p.loop.finished(p.store)

	// Wait for all the specified stores to finish.
	for _, s := range stores {
		<-p.environment[s].finished()
	}

}

// finished returns a channel that is closed when the main Done channel is closed.
func (t *Payload) finished() chan struct{} {

	// We don't want the done action to run at the same time as this function, so we use a mutex
	t.m.Lock()
	defer t.m.Unlock()

	// Create a buffered channel
	notify := make(chan struct{}, 1)

	// If the tracker has already finished, we close the notify channel before returning it.
	if t.complete {
		close(notify)
		return notify
	}

	// If the tracker hasn't finished yet, we append the notify channel to the subscribers before
	// returning it.
	t.subscribers = append(t.subscribers, notify)
	return notify

}
