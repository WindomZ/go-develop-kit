package encoding

import "testing"

func TestEncodeHash(t *testing.T) {
	if !ValidHash(EncodeHash("This is test!")) {
		t.Fatal("Fail to encode hash")
	}
}

func TestEncodeHash256(t *testing.T) {
	if !ValidHash256(EncodeHash256("This is test!")) {
		t.Fatal("Fail to encode hash")
	}
}

func TestEncodeHash512(t *testing.T) {
	if !ValidHash512(EncodeHash512("This is test!")) {
		t.Fatal("Fail to encode hash")
	}
}
