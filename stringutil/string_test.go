package stringutil

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestSubstring(t *testing.T) {
	assert.Equal(t, Substring("abc", -1, -1), "abc")
	assert.Equal(t, Substring("abc", -1, 1), "a")
	assert.Equal(t, Substring("abc", -1, 2), "ab")
	assert.Equal(t, Substring("abc", -1, 3), "abc")
	assert.Equal(t, Substring("abc", -1, 4), "abc")
	assert.Equal(t, Substring("abc", 0, 2), "ab")
	assert.Equal(t, Substring("abc", 0, 3), "abc")
	assert.Equal(t, Substring("abc", 0, 4), "abc")
	assert.Equal(t, Substring("abc", 1, 4), "bc")
	assert.Equal(t, Substring("abc", 2, 1), "c")
	assert.Equal(t, Substring("abc", 2, 2), "")
	assert.Equal(t, Substring("abc", 3, 2), "")
	assert.Equal(t, Substring("abc", 4, 2), "")
}
