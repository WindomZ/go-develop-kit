package encoding

import "testing"

func TestEncodeMD5(t *testing.T) {
	if !ValidMD5(EncodeMD5("This is test!")) {
		t.Fatal("Fail to encode md5")
	}
}
