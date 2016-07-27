package math

import "testing"

func TestFloatPrecision(t *testing.T) {
	var f float64 = 2.012345
	if r := FloatPrecision(f, 5, true); r != 2.01235 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatRound(t *testing.T) {
	var f float64 = 2.012345
	if r := FloatRound(f, 5); r != 2.01235 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatFixed(t *testing.T) {
	var f float64 = 2.012345
	if r := FloatFixed(f, 5); r != 2.01235 {
		t.Fatalf("Error %#v", r)
	}
	f = 4.012345
	if r := FloatFixed(f, 5); r != 4.01235 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatFixedToInt(t *testing.T) {
	var f float64 = 2.012345
	if r := FloatFixedToInt(f, 5); r != 201235 {
		t.Fatalf("Error %#v", r)
	}
	f = 4.012345
	if r := FloatFixedToInt(f, 5); r != 401235 {
		t.Fatalf("Error %#v", r)
	}
}
