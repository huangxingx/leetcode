package hammingDistance

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, 2, hammingDistance(1, 4))
	assert.Equal(t, 2, hammingDistance(2, 4))
}
