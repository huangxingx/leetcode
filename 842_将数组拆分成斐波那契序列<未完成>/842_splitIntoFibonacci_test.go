package splitIntoFibonacci

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSplitIntoFibonacci(t *testing.T) {
	assert.Equal(t, []int{123, 456, 579}, splitIntoFibonacci("123456579"))
	assert.Equal(t, []int{1, 1, 2, 3, 5, 8, 13}, splitIntoFibonacci("11235813"))
	assert.Equal(t, []int{}, splitIntoFibonacci("112358130"))
	assert.Equal(t, []int{}, splitIntoFibonacci("0123"))
	assert.Equal(t, []int{11, 0, 11, 11}, splitIntoFibonacci("1101111"))

}
