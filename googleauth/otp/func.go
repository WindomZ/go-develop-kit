package otp

import "github.com/WindomZ/go-develop-kit/googleauth"

func GenerateSecret(lens ...int) string {
	if lens != nil && len(lens) != 0 {
		return googleauth.GenerateRandomSecret(lens[0], false)
	}
	return googleauth.GenerateRandomSecret(googleauth.DefaultRandomSecretLength, false)
}

func ValidSecret(secret string, lens ...int) bool {
	if lens != nil && len(lens) != 0 {
		return googleauth.ValidSecret(secret, lens[0], false)
	}
	return googleauth.ValidSecret(secret, googleauth.DefaultRandomSecretLength, false)
}
