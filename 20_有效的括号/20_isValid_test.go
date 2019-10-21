package isValid

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("([)]"))
	assert.Equal(t, true, isValid("{[]}"))
	assert.Equal(t, false, isValid("]"))
	assert.Equal(t, false, isValid("[])"))
}
