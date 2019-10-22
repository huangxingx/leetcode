package defangIPaddr

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDefangIPaddr(t *testing.T) {
	assert.Equal(t, "1[.]1[.]1[.]1", defangIPaddr("1.1.1.1"))
	assert.Equal(t, "255[.]100[.]50[.]0", defangIPaddr("255.100.50.0"))
}
