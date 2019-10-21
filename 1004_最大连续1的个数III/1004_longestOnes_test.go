package longestOnes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {

	assert.Equal(t, 6, longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	assert.Equal(t, 10, longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))

}
