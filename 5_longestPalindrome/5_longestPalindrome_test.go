package longestPalindrome

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	testStringMap := map[string]string{
		"abcdeedcba": "abcdeedcba",

		"abcba": "abcba",

		"cbbd":    "bb",
		"babad":   "bab",
		"cbbb":    "bbb",
		"babcdea": "bab",
		"aa":      "aa",

		" ":  " ",
		"":   "",
		"a":  "a",
		"ac": "a",
	}

	for key, value := range testStringMap {
		assert.Equal(t, value, longestPalindrome(key))
	}
}
