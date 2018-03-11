package flux

import (
	"testing"

	"github.com/dave/ktest/assert"
)

func TestNewLoopDetector(t *testing.T) {
	s1 := &mockStore1{}
	s2 := &mockStore2{}
	d := newLoopDetector()

	found, store := d.request(s1, s2)
	assert.False(t, found)
	assert.Nil(t, store)

	found, store = d.request(s2, s1)
	assert.True(t, found)
	assert.Equal(t, s1, store)

	d.finished(s1)
	found, store = d.request(s2, s1)
	assert.False(t, found)
	assert.Nil(t, store)
}

func TestLoopDetectorInner(t *testing.T) {
	s1 := &mockStore1{}
	s2 := &mockStore2{}
	s3 := &mockStore3{}
	s4 := &mockStore4{}
	s5 := &mockStore5{}
	s6 := &mockStore6{}
	d := newLoopDetector()

	found, store := d.request(s1, s2)
	assert.False(t, found)
	assert.Nil(t, store)
	found, store = d.request(s2, s3)
	assert.False(t, found)
	assert.Nil(t, store)
	found, store = d.request(s3, s1)
	assert.True(t, found)
	assert.Equal(t, s1, store)
	found, store = d.request(s4, s5)
	assert.False(t, found)
	assert.Nil(t, store)
	found, store = d.request(s6, s4)
	assert.False(t, found)
	assert.Nil(t, store)
}

type mockStore1 struct{}
type mockStore2 struct{}
type mockStore3 struct{}
type mockStore4 struct{}
type mockStore5 struct{}
type mockStore6 struct{}

func (s *mockStore1) Handle(payload *Payload) (finished bool) { return true }
func (s *mockStore2) Handle(payload *Payload) (finished bool) { return true }
func (s *mockStore3) Handle(payload *Payload) (finished bool) { return true }
func (s *mockStore4) Handle(payload *Payload) (finished bool) { return true }
func (s *mockStore5) Handle(payload *Payload) (finished bool) { return true }
func (s *mockStore6) Handle(payload *Payload) (finished bool) { return true }
