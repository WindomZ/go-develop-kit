package encoding

import "testing"

func TestEncodeBase64String(t *testing.T) {
	const STR = "This is test!"
	if DecodeBase64String(EncodeBase64String(STR)) != STR {
		t.Fatal("Fail to encode or decode")
	}
}
