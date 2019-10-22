package flipAndInvertImage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFlipAndInvertImage(t *testing.T) {
	assert.Equal(t, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}, flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	assert.Equal(t, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0,0,0,1},{1, 0, 1, 0}}, flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))

}
