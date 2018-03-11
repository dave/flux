package mock_flux

//go:generate go get github.com/golang/mock/mockgen
//go:generate mockgen -destination mocks.go github.com/dave/flux DispatcherInterface,NotifierInterface,AppInterface,WatcherInterface
