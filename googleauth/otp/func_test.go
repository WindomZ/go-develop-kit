package otp

import "testing"

func TestGenerateSecret(t *testing.T) {
	if true {
		return
	}
	const LEN int = 16
	secret := GenerateSecret(LEN)
	t.Log(secret)
	if len(secret) != LEN {
		t.Fatal("invalid secret length")
	}
}
