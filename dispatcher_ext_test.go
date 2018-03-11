package flux_test

import (
	"testing"

	"github.com/dave/flux"
	"github.com/dave/flux/mock_flux"

	"github.com/golang/mock/gomock"
)

func TestDispatcher_Notify(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	not := mock_flux.NewMockNotifierInterface(m)
	a := &st1{}
	gomock.InOrder(not.EXPECT().Notify().Return(closedChannel()))
	d := flux.NewDispatcher(not, a)
	done := d.Dispatch(false)
	<-done
}

func closedChannel() chan struct{} {
	c := make(chan struct{}, 1)
	close(c)
	return c
}

func TestDispatcher_NotifyAfter(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	not := mock_flux.NewMockNotifierInterface(m)
	a := &st1{}
	gomock.InOrder(not.EXPECT().Notify().Return(closedChannel()))
	d := flux.NewDispatcher(not, a)
	done := d.Dispatch(true)
	<-done
}

type st1 struct{}

func (s *st1) Handle(payload *flux.Payload) (finished bool) {
	async := payload.Action.(bool)
	if async {
		go func() {
			payload.Notify()
			close(payload.Done)
		}()
		return false
	} else {
		payload.Notify()
		return true
	}
}
