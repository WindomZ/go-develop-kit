package encoding

import "encoding/base64"

// EncodeBase64 returns the base64 bytes encoding of bytes data.
func EncodeBase64(data []byte) []byte {
	result := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(result, data)
	return result
}

// DecodeBase64 returns the bytes represented by the base64 bytes data.
func DecodeBase64(data []byte) ([]byte, error) {
	result := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(result, data)
	return result[:n], err
}

// EncodeBase64ToString returns the base64 encoding of data.
func EncodeBase64ToString(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeBase64FromString returns the bytes represented by the base64 string data.
func DecodeBase64FromString(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// EncodeBase64String returns the base64 string encoding of string data.
func EncodeBase64String(data string) string {
	return EncodeBase64ToString([]byte(data))
}

// DecodeBase64String returns the string represented by the base64 string data.
func DecodeBase64String(data string) string {
	if b, err := DecodeBase64FromString(data); err == nil {
		return string(b)
	}
	return ""
}
