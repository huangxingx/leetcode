package plusOne

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert.Equal(t, []int{1, 3, 0}, plusOne([]int{1, 2, 9}))
	assert.Equal(t, []int{2, 0, 0}, plusOne([]int{1, 9, 9}))
	assert.Equal(t, []int{1, 0, 0, 0}, plusOne([]int{9, 9, 9}))
}
