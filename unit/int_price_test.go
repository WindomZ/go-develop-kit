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

func TestJSONIntPrice(t *testing.T) {
	SetIntPricePrecision(5)
	p := &testIntPrice{
		Price1: NewIntPrice(101234),
		Price2: NewIntPriceFloat(2.012345),
		Price3: NewIntPriceString("301234"),
		Price4: NewIntPriceFloatString("4.012345"),
	}
	t.Logf("%#v", p)
	data, err := gojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", string(data))
	var pp testIntPrice
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", pp)
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
