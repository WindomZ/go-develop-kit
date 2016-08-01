package unit

import (
	"github.com/WindomZ/go-macro/json"
	"testing"
)

type testIntPrice struct {
	Price1 IntPrice `json:"price1"`
	Price2 IntPrice `json:"price2"`
	Price3 IntPrice `json:"price3"`
	Price4 IntPrice `json:"price4"`
}

func TestNewIntPrice(t *testing.T) {
	SetIntPricePrecision(5)
}

func TestJSONIntPrice(t *testing.T) {
	p := &testIntPrice{
		Price1: NewIntPrice(101234),
		Price2: NewIntPriceFloat(2.012345),
		Price3: NewIntPriceString("301234"),
		Price4: NewIntPriceFloatString("4.012345"),
	}
	data, err := gojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	var pp testIntPrice
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	if p.Price1.Float64() != 1.01234 {
		t.Fatal("Error Price1:", p.Price1.Float64())
	}
	if p.Price2.Float64() != 2.01235 {
		t.Fatal("Error Price2:", p.Price2.Float64())
	}
	if p.Price3.Float64() != 3.01234 {
		t.Fatal("Error Price3:", p.Price3.Float64())
	}
	if p.Price4.Float64() != 4.01235 {
		t.Fatal("Error Price4:", p.Price4.Float64())
	}
}

func TestIntPrice_Float64(t *testing.T) {
	p := NewIntPriceFloat(1.012345)
	p.SetFloat64(2.0125, 3)
	if p.Float64() != 2.013 {
		t.Fatal("Error:", p.Float64())
	}
	p.SetFloat64(4.0125, 3)
	if p.Float64() != 4.013 {
		t.Fatal("Error:", p.Float64())
	}
}

func TestIntPrice_SetFloat64(t *testing.T) {
	p := NewIntPriceFloat(1.012345)
	p.SetFloat64(2.0125, 3)
	if p.Int64() != 201300 {
		t.Fatal("Error:", p.Int64())
	}
	p.SetFloat64(4.0125, 3)
	if p.Int64() != 401300 {
		t.Fatal("Error:", p.Int64())
	}
}

func TestIntPrice_ReciprocalFloat64(t *testing.T) {
	p := NewIntPriceFloat(1.012345)
	if p.ReciprocalFloat64() != 0.9878 {
		t.Fatal("Error:", p.ReciprocalFloat64())
	}
}

func TestIntPrice_GetMul(t *testing.T) {
	p1 := NewIntPrice(101234)
	p2 := NewIntPriceFloat(2.012345)
	p3 := NewIntPriceFloat(3.012345)
	p := p1.GetMul(p2, p3)
	if p.Int64() != 613669 {
		t.Fatal(p.Int64())
	}
}

func TestIntPrice_GetQuo(t *testing.T) {
	p1 := NewIntPrice(501234)
	p2 := NewIntPriceFloat(2.012345)
	p := p1.GetQuo(p2)
	if p.Int64() != 2 {
		t.Fatal(p.Int64())
	}
}

func TestIntPrice_GetNegation(t *testing.T) {
	var i int64 = 0
	for ; i < 1000000; i++ {
		p := NewIntPrice(i)
		if p.Int64() != i || p.Int64() != -p.GetNegation().Int64() {
			t.Fatalf("%v:%v:%v:%v", i, p.Int64(), p.GetNegation().Int64(), -p.GetNegation().Int64())
		}
	}
}
