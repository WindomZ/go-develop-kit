package math

import "testing"

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

func TestFloatSumFixed(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatSumFixed(F1, F2, 5, F3, F4, F5); r != 15.06173 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatSubFixed(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatSubFixed(F1, F2, 5, F3, F4, F5); r != -13.03704 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatMulFixed(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatMulFixed(F1, F2, 5, F3, F4, F5); r != 123.41698 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatDivFixed(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatDivFixed(F1, F2, 5, F3, F4, F5); r != 0.0083 {
		t.Fatalf("Error %#v", r)
	}
}
