package math

import "testing"

func TestFloatRound(t *testing.T) {
	var f float64 = 2.012345
	if r := FloatRound(f, 5); r != 2.01235 {
		t.Fatalf("Error %#v", r)
	}
	f = 4.012345
	if r := FloatRound(f, 5); r != 4.01235 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatRoundToInt(t *testing.T) {
	var f float64 = 2.012345
	if r := FloatRoundToInt(f, 5); r != 201235 {
		t.Fatalf("Error %#v", r)
	}
	f = 4.012345
	if r := FloatRoundToInt(f, 5); r != 401235 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatSumRound(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatSum(F1, F2); r != 3.02469 {
		t.Fatalf("Error %#v", r)
	} else if r := FloatSumRound(F1, F2, 5, F3, F4, F5); r != 15.06173 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatSubRound(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatSub(F1, F2); r != -1 {
		t.Fatalf("Error %#v", r)
	} else if r := FloatSubRound(F1, F2, 5, F3, F4, F5); r != -13.03704 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatMulRound(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatMul(F1, F2); r != 2.037187399025 {
		t.Fatalf("Error %#v", r)
	} else if r := FloatMulRound(F1, F2, 5, F3, F4, F5); r != 123.41698 {
		t.Fatalf("Error %#v", r)
	}
}

func TestFloatDivRound(t *testing.T) {
	const F1 float64 = 1.012345
	const F2 float64 = 2.012345
	const F3 float64 = 3.012345
	const F4 float64 = 4.012345
	const F5 float64 = 5.012345
	if r := FloatDiv(F1, F2); r != 0.5030673169859045 {
		t.Fatalf("Error %#v", r)
	} else if r := FloatDivRound(F1, F2, 5, F3, F4, F5); r != 0.0083 {
		t.Fatalf("Error %#v", r)
	}
}
