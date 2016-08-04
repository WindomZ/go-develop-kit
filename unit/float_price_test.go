package unit

import (
	"github.com/WindomZ/go-macro/json"
	"testing"
)

type testFloatPrice struct {
	Price1 FloatPrice `json:"price1"`
	Price2 FloatPrice `json:"price2"`
	Price3 FloatPrice `json:"price3"`
	Price4 FloatPrice `json:"price4"`
}

func TestNewFloatPrice(t *testing.T) {
	SetFloatPricePrecision(5)
}

func TestJSONFloatPrice(t *testing.T) {
	p := &testFloatPrice{
		Price1: NewFloatPrice(1.012345),
		Price2: NewFloatPriceInt(201234),
		Price3: NewFloatPriceString("3.012345"),
		Price4: NewFloatPriceIntString("401234"),
	}
	data, err := gojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	var pp testFloatPrice
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	if p.Price1.Float64() != 1.01235 {
		t.Fatal("Error Price1:", p.Price1.Float64())
	}
	if p.Price2.Float64() != 2.01234 {
		t.Fatal("Error Price2:", p.Price2.Float64())
	}
	if p.Price3.Float64() != 3.01235 {
		t.Fatal("Error Price3:", p.Price3.Float64())
	}
	if p.Price4.Float64() != 4.01234 {
		t.Fatal("Error Price4:", p.Price4.Float64())
	}
}

func TestFloatPrice_Int64(t *testing.T) {
	p := NewFloatPrice(1.012345)
	if p.Int64() != 101235 {
		t.Fatal("Error:", p.Int64())
	}
}

func TestFloatPrice_SetInt64(t *testing.T) {
	p := NewFloatPrice(1.012345)
	if p.Int64() != 101235 {
		t.Fatal("Error:", p.Int64())
	}
	p.SetInt64(201235)
	if p.Int64() != 201235 {
		t.Fatal("Error:", p.Int64())
	}
}

func TestFloatPrice_ReciprocalFloat64(t *testing.T) {
	p := NewFloatPrice(1.012345)
	if p.ReciprocalFloat64() != 0.9878 {
		t.Fatal("Error:", p.ReciprocalFloat64())
	}
}

func TestFloatPrice_Round(t *testing.T) {
	p := NewFloatPrice(1.0125)
	if p.Round(3); p.Float64() != 1.013 {
		t.Fatal("Error:", p.Float64())
	}
}

func TestFloatPrice_Add(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	if p1.Add(p2).Float64() != 3.0247 {
		t.Fatal("Error:", p1.Float64())
	}
}

func TestFloatPrice_Sub(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	if p1.Sub(p2).Float64() != -1 {
		t.Fatal("Error:", p1.Float64())
	}
}

func TestFloatPrice_Mul(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	if p1.Mul(p2).Float64() != 2.0372 {
		t.Fatal("Error:", p1.Float64())
	}
}

func TestFloatPrice_Quo(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	if p1.Quo(p2).Float64() != 0.50307 {
		t.Fatal("Error:", p1.Float64())
	}
}

func TestFloatPrice_GetSum(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	p3 := NewFloatPrice(3.012345)
	if p := p1.GetSum(p2, p3); p.Float64() != 6.03705 {
		t.Fatal("Error:", p.Float64())
	}
}

func TestFloatPrice_GetDiff(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	p3 := NewFloatPrice(3.012345)
	if p := p1.GetDiff(p2, p3); p.Float64() != -4.01235 {
		t.Fatal("Error:", p.Float64())
	}
}

func TestFloatPrice_GetMul(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	p3 := NewFloatPrice(3.012345)
	if p := p1.GetMul(p2, p3); p.Float64() != 6.13677 {
		t.Fatal("Error:", p.Float64())
	}
}

func TestFloatPrice_GetDiv(t *testing.T) {
	p1 := NewFloatPrice(1.012345)
	p2 := NewFloatPrice(2.012345)
	p3 := NewFloatPrice(3.012345)
	if p := p1.GetDiv(p2, p3); p.Float64() != 0.167 {
		t.Fatal("Error:", p.Float64())
	}
}

func TestFloatPrice_GetNegation(t *testing.T) {
	var i int64 = 0
	for ; i < 10; i++ {
		p := NewFloatPriceInt(i)
		if p.Int64() != i || p.Float64() != -p.GetNegation().Float64() {
			t.Fatalf("%v:%v:%v:%v:%v", i, p.Int64(), p.Float64(), p.GetNegation().Float64(), -p.GetNegation().Float64())
		}
	}
}
