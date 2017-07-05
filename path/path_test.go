package path

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestIsExist(t *testing.T) {

	ok, err := IsExist(ExecPath())
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestEnsure(t *testing.T) {
	assert.NoError(t, Ensure(ExecPath(), false))
	assert.NoError(t, Ensure(ExecDir(), true))
}

func TestExecPath(t *testing.T) {
	assert.NotEmpty(t, ExecPath())
}

func TestExecDir(t *testing.T) {
	assert.NotEmpty(t, ExecDir())
}
