package stringutil

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestRegexpSubstring(t *testing.T) {
	assert.Equal(t, RegexpSubstring(`{.+}`, "{abc}", "{", "}"), []string{"abc"})
	assert.Equal(t, RegexpSubstring(`s{.+}`, "{abc}", "{", "}"), []string(nil))
	assert.Equal(t, RegexpSubstring(`{[^{}]+}`, "{abc}, {def}", "{", "}"), []string{"abc", "def"})
}
