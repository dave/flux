package flux

import (
	"testing"

	"github.com/dave/ktest/assert"
)

func TestNotifier_Watch(t *testing.T) {

	n := NewNotifier()

	var a, b bool
	n.Watch("a", func(done chan struct{}) { a = true; close(done) })
	n.Watch("b", func(done chan struct{}) { b = true; close(done) })
	done := n.Notify()
	<-done
	assert.True(t, a)
	assert.True(t, b)

	var c, d bool
	n.Watch("c", func(done chan struct{}) { c = true; close(done) })
	n.Watch("d", func(done chan struct{}) { d = true; close(done) })
	n.Delete("c")
	done = n.Notify()
	<-done
	assert.False(t, c)
	assert.True(t, d)

}
