package math

import "testing"

func TestMaxInt64(t *testing.T) {
	if MaxInt64(64, 16) != 64 {
		t.Fatal("Error MaxInt", MaxInt64(64, 16))
	}
}

func TestMinInt64(t *testing.T) {
	if MinInt64(64, 16) != 16 {
		t.Fatal("Error MaxInt", MinInt64(64, 16))
	}
}
