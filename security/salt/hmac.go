package salt

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func AddHMACMD5Salt(salt, value string) string {
	mac := hmac.New(md5.New, []byte(salt))
	mac.Write([]byte(value))
	r := mac.Sum(nil)
	return hex.EncodeToString(r[:])
}

func VerifyHMACMD5Salt(salt, value, reference string) bool {
	if len(salt) == 0 || len(value) == 0 || len(reference) == 0 {
		return false
	}
	return strings.EqualFold(strings.ToLower(reference), strings.ToLower(AddHMACMD5Salt(salt, value)))
}
