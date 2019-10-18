package climbStairs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, 2, climbStairs(2))
	assert.Equal(t, 1, climbStairs(1))
	assert.Equal(t, 3, climbStairs(3))
}
