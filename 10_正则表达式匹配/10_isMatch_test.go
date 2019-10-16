package isMath


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("aa", "a."))
	assert.Equal(t, true, isMatch("aa", "a*"))

	assert.Equal(t, true, isMatch("ab", ".*"))

}