package isHappy

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	assert.Equal(t, true, isHappy(19))

}
