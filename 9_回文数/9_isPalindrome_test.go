package isPalindrome

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestMyAtoi(t *testing.T) {
	testStringMap := map[int]bool{
		121: true,

		123: false,

		-123: false,

	}

	for key, value := range testStringMap {
		assert.Equal(t, value, isPalindrome2(key), fmt.Sprintf("key is %d", key))
	}
}
