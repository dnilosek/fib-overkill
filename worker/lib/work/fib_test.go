package work_test

import (
	"testing"

	. "github.com/dnilosek/fib-overkill/worker/lib/work"
	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {

	assert.Equal(t, 0, FibonacciNumberAtIndex(0))
	assert.Equal(t, 1, FibonacciNumberAtIndex(1))
	assert.Equal(t, 1, FibonacciNumberAtIndex(2))
	assert.Equal(t, 2, FibonacciNumberAtIndex(3))
	assert.Equal(t, 3, FibonacciNumberAtIndex(4))
	assert.Equal(t, 5, FibonacciNumberAtIndex(5))
	assert.Equal(t, 8, FibonacciNumberAtIndex(6))
}
