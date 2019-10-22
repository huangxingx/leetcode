package canWinNim

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCanWinNim(t *testing.T) {
	assert.Equal(t, false, canWinNim(4))
}
