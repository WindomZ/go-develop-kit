package googleauth

import (
	"encoding/base32"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomSecret(t *testing.T) {
	for i := 0; i < 1000; i++ {
		secret := GenerateRandomSecret(20, true)
		_, err := base32.StdEncoding.DecodeString(secret)
		assert.Nil(t, err)
	}
}
