package stringutil

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestRegexpSubstrings(t *testing.T) {
	assert.Equal(t, RegexpSubstrings(`{.+}`, "{abc}", "{", "}"), []string{"abc"})
	assert.Equal(t, RegexpSubstrings(`s{.+}`, "{abc}", "{", "}"), []string(nil))
	assert.Equal(t, RegexpSubstrings(`{[^{}]+}`, "{abc}, {def}", "{", "}"), []string{"abc", "def"})
}

func TestRegexpSubstring(t *testing.T) {
	assert.Equal(t, RegexpSubstring(`{.+}`, "{abc}", "{", "}"), "abc")
	assert.Equal(t, RegexpSubstring(`s{.+}`, "{abc}", "{", "}"), "")
	assert.Equal(t, RegexpSubstring(`{[^{}]+}`, "{abc}, {def}", "{", "}"), "abc")
}
