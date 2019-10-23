package maxProfit2

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 7, maxProfit([]int{7,1,5,3,6,4}))
	assert.Equal(t, 4, maxProfit([]int{1,2,3,4,5}))
}
