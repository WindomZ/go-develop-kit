package otp

import "github.com/WindomZ/go-develop-kit/googleauth"

func GenerateSecret() string {
	return googleauth.GenerateRandomSecret(googleauth.DefaultRandomSecretLength, true)
}

func ValidSecret(secret string) bool {
	return googleauth.ValidSecret(secret, googleauth.DefaultRandomSecretLength, true)
}
