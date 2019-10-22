package rob

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	//assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 4, rob([]int{2, 1, 1, 2}))
}
