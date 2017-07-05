package path

import (
	"path"
	"testing"

	"github.com/WindomZ/testify/assert"
)

var testFilePath = ""

func TestCreateFile(t *testing.T) {
	testFilePath = path.Join(ExecDir(), "test")
	if MustExist(testFilePath) {
		assert.NoError(t, RemoveFile(testFilePath, false))
	}
	assert.NoError(t, CreateFile(testFilePath))
}

func TestAppendToFile(t *testing.T) {
	if MustExist(testFilePath) {
		assert.NoError(t, RemoveFile(testFilePath, false))
	}

	assert.NoError(t, AppendToFile(testFilePath, "Hello", "你好"))
	s, err := ReadFile(testFilePath)
	assert.NoError(t, err)
	assert.Equal(t, s, `Hello
你好
`)

	assert.NoError(t, AppendToFile(testFilePath, "World", "世界"))
	s, err = ReadFile(testFilePath)
	assert.NoError(t, err)
	assert.Equal(t, s, `Hello
你好
World
世界
`)
}

func TestOverwriteFile(t *testing.T) {
	if MustExist(testFilePath) {
		assert.NoError(t, OverwriteFile(testFilePath, "Hello World", "你好世界"))
	}

	s, err := ReadFile(testFilePath)
	assert.NoError(t, err)
	assert.Equal(t, s, `Hello World
你好世界
`)
}

func TestRemoveFile(t *testing.T) {
	assert.NoError(t, RemoveFile(testFilePath, false))
}
