package convert

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, "LCIRETOESIIGEDHN", convert("LEETCODEISHIRING", 3))
	assert.Equal(t, "LDREOEIIECIHNTSG", convert("LEETCODEISHIRING", 4))

}
