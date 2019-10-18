package fib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 1, fib(1))
	assert.Equal(t, 1, fib(2))
	assert.Equal(t, 2, fib(3))
}
