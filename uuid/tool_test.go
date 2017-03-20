package uuid

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestToUUID(t *testing.T) {
	assert.Equal(t, ToUUID("FDB084e4-fe34-5c42-9016-15819df8be03"),
		"fdb084e4-fe34-5c42-9016-15819df8be03")
}

func TestNoDashUUID(t *testing.T) {
	assert.Equal(t, NoDashUUID("FDB084e4-fe34-5c42-9016-15819df8be03"),
		"fdb084e4fe345c42901615819df8be03")
}
