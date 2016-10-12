package encoding

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// EncodeHash returns the sha1 encoding of s.
func EncodeHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// ValidHash returns true if length of s same as size of a SHA1 checksum.
func ValidHash(s string) bool {
	return len(s) == sha1.Size*2
}

// EncodeHash256 returns the sha256 encoding of s.
func EncodeHash256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// ValidHash256 returns true if length of s same as size of a SHA256 checksum.
func ValidHash256(s string) bool {
	return len(s) == sha256.Size*2
}

// EncodeHash512 returns the sha256 encoding of s.
func EncodeHash512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// ValidHash512 returns true if length of s same as size of a SHA-512 checksum.
func ValidHash512(s string) bool {
	return len(s) == sha512.Size*2
}
