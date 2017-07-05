package path

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestHomeDir(t *testing.T) {
	home, err := HomeDir()
	assert.NoError(t, err)
	assert.NotEmpty(t, home)
}
