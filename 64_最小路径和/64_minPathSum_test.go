package minPathSum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	assert.Equal(t, 7, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))

}
