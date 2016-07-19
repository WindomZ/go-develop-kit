package googleauth

import (
	"encoding/base32"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomSecret(t *testing.T) {
	// Brute force test to check if the returned string is
	// a valid Base32 string.
	for i := 0; i < 1000; i++ {
		secret := generateRandomSecret(20, true)
		_, err := base32.StdEncoding.DecodeString(secret)
		assert.Nil(t, err)
	}
}
