package encoding

import (
	"bytes"
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	var DATA = []byte("This is test!")
	if data, err := DecodeBase64(EncodeBase64(DATA)); err != nil {
		t.Fatal(err)
	} else if !bytes.Equal(data, DATA) {
		t.Fatal("Fail to encode or decode")
	}
}

func TestEncodeBase64ToString(t *testing.T) {
	var DATA = []byte("This is test!")
	if data, err := DecodeBase64FromString(EncodeBase64ToString(DATA)); err != nil {
		t.Fatal(err)
	} else if !bytes.Equal(data, DATA) {
		t.Fatal("Fail to encode or decode")
	}
}

func TestEncodeBase64String(t *testing.T) {
	const STR = "This is test!"
	if DecodeBase64String(EncodeBase64String(STR)) != STR {
		t.Fatal("Fail to encode or decode")
	}
}
