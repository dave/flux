package flux

type ActionInterface interface{}

type AppInterface interface {
	WatcherInterface
	DispatcherInterface
}

type DispatcherInterface interface {
	Dispatch(action ActionInterface) chan struct{}
}

type WatcherInterface interface {
	Watch(key interface{}, f func(done chan struct{}))
	Delete(key interface{})
}

type NotifierInterface interface {
	// Notify sends the notif notification to all subscribers of that
	// notification for object. If object is nil, the notification is sent to
	// all subscribers. The chanel returned is closed when the notify action
	// has finished.
	Notify() (done chan struct{})
}
