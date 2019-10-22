package maxSubArray

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
