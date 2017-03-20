package uuid

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestNewUUID(t *testing.T) {
	assert.Equal(t, Valid(NewUUID()), true)
}

func TestNewUUIDWithName(t *testing.T) {
	assert.Equal(t, Valid(NewUUIDWithName("hello")), true)
}

func TestNewSafeUUID(t *testing.T) {
	for i := 0; i < 10000; i++ {
		assert.Equal(t, Valid(NewSafeUUID()), true)
	}
}
