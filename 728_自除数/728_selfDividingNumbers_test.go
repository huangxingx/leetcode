package selfDividingNumbers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSelfDividingNumbers(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}, selfDividingNumbers(1, 22))

}
