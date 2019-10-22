package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	assert.Equal(t, 3, game([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.Equal(t, 2, game([]int{2, 2, 3}, []int{1, 2, 3}))
}
