package googleauth

import (
	"crypto/rand"
	"encoding/base32"
)

const alphanum string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateRandomSecret(size int, encodeToBase32 bool) string {
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
