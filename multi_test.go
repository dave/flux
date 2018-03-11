package flux

import (
	"testing"

	"github.com/dave/ktest/assert"
)

func TestWatchMulti(t *testing.T) {
	done := make(chan struct{})
	go func() {
		// We watch two channels, a and b
		a := make(chan struct{}, 1)
		b := make(chan struct{}, 1)
		c := Multi(a, b)

		// We send a signal to a and wait for the response on the
		a <- struct{}{}
		assert.WaitFor(t, c, true, "A")

		// We send a signal to a and b and ensure two signals are returned
		a <- struct{}{}
		b <- struct{}{}
		assert.WaitFor(t, c, true, "B")
		assert.WaitFor(t, c, true, "C")

		// We close a and send on b to ensure it still works
		close(a)
		b <- struct{}{}
		assert.WaitFor(t, c, true, "D")

		// We close the second input channel, and check that the signal channel is also closed
		close(b)
		assert.WaitFor(t, c, false, "E")

		close(done)
	}()
	assert.WaitFor(t, done, false, "G")
}
