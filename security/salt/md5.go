package salt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func AddMD5Salt(salt, value string) string {
	r := md5.Sum([]byte(salt + value))
	return hex.EncodeToString(r[:])
}

func VerifyMD5Salt(salt, value, reference string) bool {
	if len(salt) == 0 || len(value) == 0 || len(reference) == 0 {
		return false
	}
	return strings.EqualFold(strings.ToLower(reference), strings.ToLower(AddMD5Salt(salt, value)))
}
