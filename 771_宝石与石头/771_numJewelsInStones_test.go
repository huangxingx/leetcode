package numJewelsInStones

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNumJewelsInStones(t *testing.T) {
	assert.Equal(t, 3, numJewelsInStones("aA", "aAAbbbb"))
	assert.Equal(t, 0, numJewelsInStones("z", "ZZ"))

}