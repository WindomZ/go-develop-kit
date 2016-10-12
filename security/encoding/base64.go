package encoding

import "encoding/base64"

// EncodeBase64String returns the base64 encoding of data.
func EncodeBase64String(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// DecodeBase64String returns the bytes represented by the base64 string data.
// if error returns empty string.
func DecodeBase64String(data string) string {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(b)
}
