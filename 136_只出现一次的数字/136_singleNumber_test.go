package singleNumber

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
	assert.Equal(t, 4, singleNumber([]int{4,1,2,1,2}))

}
