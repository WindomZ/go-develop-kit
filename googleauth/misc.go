package googleauth

import (
	"crypto/rand"
	"encoding/base32"
)

const alphanum string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const alphanumBase32 string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func GenerateRandomSecret(size int, encodeToBase32 bool) string {
	var bytes = make([]byte, size)
	rand.Read(bytes)
	if encodeToBase32 {
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		return base32.StdEncoding.EncodeToString(bytes)
	} else {
		for i, b := range bytes {
			bytes[i] = alphanumBase32[b%byte(len(alphanumBase32))]
		}
	}
	return string(bytes)
}

func ValidSecret(secret string, size int, encodeToBase32 bool) bool {
	if encodeToBase32 {
		if _, err := base32.StdEncoding.DecodeString(secret); err != nil {
			return false
		}
	} else if len(secret) != size {
		return false
	} else {
		//TODO: check alphanumBase32
	}
	return true
}
