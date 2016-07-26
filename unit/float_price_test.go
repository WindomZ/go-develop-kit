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

func TestJSONFloatPrice(t *testing.T) {
	SetFloatPricePrecision(5)
	p := &testFloatPrice{
		Price1: NewFloatPrice(1.012345),
		Price2: NewFloatPriceInt(201234),
		Price3: NewFloatPriceString("3.012345"),
		Price4: NewFloatPriceIntString("401234"),
	}
	t.Logf("%#v", p)
	data, err := gojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", string(data))
	var pp testFloatPrice
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", pp)
}

func TestFloatPrice_GetNegation(t *testing.T) {
	SetFloatPricePrecision(8)
	var i int64 = 0
	for ; i < 1000000; i++ {
		p := NewFloatPriceInt(i)
		if p.Int64() != i || p.Float64() != -p.GetNegation().Float64() {
			t.Fatalf("%v:%v:%v:%v:%v", i, p.Int64(), p.Float64(), p.GetNegation().Float64(), -p.GetNegation().Float64())
		}
	}
}
