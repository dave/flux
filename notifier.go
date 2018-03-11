package flux

import "sync"

type Notifier struct {
	m           sync.RWMutex
	subscribers map[interface{}]func(done chan struct{})
}

func NewNotifier() *Notifier {
	n := &Notifier{}
	n.subscribers = map[interface{}]func(done chan struct{}){}
	return n
}

func (s *Notifier) Delete(key interface{}) {
	s.delete(key)
}

func (s *Notifier) Watch(key interface{}, f func(done chan struct{})) {
	s.set(key, f)
}

func (s *Notifier) Notify() (done chan struct{}) {
	wg := &sync.WaitGroup{}
	all := s.values()
	wg.Add(len(all))
	for _, f := range all {
		finished := make(chan struct{}, 1)
		go func() {
			<-finished
			wg.Done()
		}()
		f(finished)
	}
	done = make(chan struct{}, 1)
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}

func (n *Notifier) set(key interface{}, value func(done chan struct{})) {
	n.m.Lock()
	defer n.m.Unlock()
	n.subscribers[key] = value
}

func (n *Notifier) delete(key interface{}) {
	n.m.Lock()
	defer n.m.Unlock()
	delete(n.subscribers, key)
}

func (n *Notifier) values() []func(done chan struct{}) {
	n.m.RLock()
	defer n.m.RUnlock()
	var out []func(done chan struct{})
	for _, v := range n.subscribers {
		out = append(out, v)
	}
	return out
}
