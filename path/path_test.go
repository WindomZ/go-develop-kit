package path

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestIsExist(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	filePath, _ := filepath.Abs(file)

	ok, err := IsExist(filePath)
	assert.NoError(t, err)
	assert.True(t, ok)
}
