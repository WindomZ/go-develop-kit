package encoding

import (
	"crypto/md5"
	"encoding/hex"
)

const EmptyMD5Hash string = "00000000000000000000000000000000"

// EncodeMD5 returns the md5 encoding of s.
func EncodeMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// ValidMD5 returns true if length of s same as size of an MD5 checksum.
func ValidMD5(s string) bool {
	return len(s) == md5.Size*2
}
