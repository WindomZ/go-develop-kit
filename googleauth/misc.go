package googleauth

import (
	"crypto/rand"
	"encoding/base32"
)

func generateRandomSecret(size int, encodeToBase32 bool) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	if encodeToBase32 {
		return base32.StdEncoding.EncodeToString(bytes)
	}
	return string(bytes)
}
