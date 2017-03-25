package uuid

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, IsEmpty("00000000-0000-0000-0000-000000000000"), true)
	assert.Equal(t, IsEmpty("00000000-0000-0000-0000-000000000001"), false)
}

func TestValid(t *testing.T) {
	assert.Equal(t, Valid("00000000-0000-0000-0000-000000000000"), true)
	assert.Equal(t, Valid("00000000-0000-0000-0000-000000000001"), true)

	assert.Equal(t, Valid("000000000000-0000-0000-000000000001"), false)
	assert.Equal(t, Valid("00000000000000000000000000000001"), true)
}

func TestValidNoEmpty(t *testing.T) {
	assert.Equal(t, ValidNoEmpty("00000000-0000-0000-0000-000000000000"), false)
	assert.Equal(t, ValidNoEmpty("00000000-0000-0000-0000-000000000001"), true)

	assert.Equal(t, ValidNoEmpty("00000000000000000000000000000000"), false)
	assert.Equal(t, ValidNoEmpty("00000000000000000000000000000001"), true)
}

func TestValidNoDash(t *testing.T) {
	assert.Equal(t, ValidNoDash("00000000-0000-0000-0000-000000000001"), false)
	assert.Equal(t, ValidNoDash("00000000000000000000000000000001"), true)
}
