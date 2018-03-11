package flux

import (
	"testing"

	"github.com/dave/ktest/assert"
)

type st struct{}

func (s *st) Handle(payload *Payload) (finished bool) { return true }

func TestNewDispatcher(t *testing.T) {
	a := &st{}
	b := &st{}
	c := NewDispatcher(nil, a, b)
	assert.Equal(t, 2, len(c.stores))
	assert.Equal(t, a, c.stores[0])
	assert.Equal(t, b, c.stores[1])
}

func TestDispatcher_Register(t *testing.T) {
	a := &st{}
	b := NewDispatcher(nil)
	b.Register(a)
	assert.Equal(t, 1, len(b.stores))
	assert.Equal(t, a, b.stores[0])
}

func TestDispatcher_Dispatch(t *testing.T) {
	a := &st1{}
	b := &st1{}
	c := &st1{}
	d := NewDispatcher(nil, a, b)
	d.Register(c)
	done := d.Dispatch("e")
	<-done
	assert.Equal(t, "e", a.handled)
	assert.Equal(t, "e", b.handled)
	assert.Equal(t, "e", c.handled)
}

type st1 struct {
	handled ActionInterface
}

func (s *st1) Handle(payload *Payload) (finished bool) {
	s.handled = payload.Action
	return true
}
