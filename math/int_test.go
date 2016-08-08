package math

import "testing"

func TestMaxInt(t *testing.T) {
	if MaxInt(64, 16) != 64 {
		t.Fatal("Error MaxInt", MaxInt(64, 16))
	}
}

func TestMinInt(t *testing.T) {
	if MinInt(64, 16) != 16 {
		t.Fatal("Error MinInt", MinInt(64, 16))
	}
}
