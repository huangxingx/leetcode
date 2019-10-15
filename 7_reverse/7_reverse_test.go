package reverse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testStringMap := map[int]int{
		12:  21,
		123: 321,
		100: 1,
		120: 21,
		-120: -21,
		1534236469: 0,
	}

	for key, value := range testStringMap {
		assert.Equal(t, value, reverse(key))
	}
}
