package myAtoi

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	testStringMap := map[string]int{
		"  qw-123": 0,
		"  -123wd": -123,
		"  -123wd   ": -123,
		"  w+123wd   ": 0,

		"123":    123,
		"  123":  123,
		"  -123": -123,

		"  -qw-123": 0,
		"-qw-123":   0,
		"qw-123":    0,

		"-91283472332": -2147483648,
	}

	for key, value := range testStringMap {
		assert.Equal(t, value, myAtoi(key))
	}
}
