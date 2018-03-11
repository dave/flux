package flux

import "sync"

func newLoopDetector() *loopDetector {
	return &loopDetector{
		waiting: map[StoreInterface][]StoreInterface{},
	}
}

type loopDetector struct {
	m       sync.RWMutex
	waiting map[StoreInterface][]StoreInterface
}

// finished marks a store as finished processing
func (w *loopDetector) finished(store StoreInterface) {
	w.delete(store)
}

// request requests a wait. We return if a loop was not found, and if one is found we return the store
// that was waiting for this store.
func (w *loopDetector) request(store StoreInterface, waitFor ...StoreInterface) (loopFound bool, loopStore StoreInterface) {
	// returns true if s1 is waiting for s2
	var isWaiting func(s1, s2 StoreInterface) bool
	isWaiting = func(s1, s2 StoreInterface) bool {
		waits, ok := w.get(s1)
		if !ok {
			return false
		}
		for _, inner := range waits {
			if inner == s2 {
				return true
			}
			if isWaiting(inner, s2) {
				return true
			}
		}
		return false
	}
	for _, requested := range waitFor {
		if isWaiting(requested, store) {
			return true, requested
		}
	}
	w.set(store, waitFor)
	return false, nil
}

func (w *loopDetector) get(k StoreInterface) ([]StoreInterface, bool) {
	w.m.RLock()
	defer w.m.RUnlock()
	v, ok := w.waiting[k]
	return v, ok
}

func (w *loopDetector) set(k StoreInterface, v []StoreInterface) {
	w.m.Lock()
	defer w.m.Unlock()
	w.waiting[k] = v
}

func (w *loopDetector) delete(k StoreInterface) {
	w.m.Lock()
	defer w.m.Unlock()
	delete(w.waiting, k)
}
