package unit

import (
	"github.com/WindomZ/go-macro/json"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	SetCurrencyMapping("123", "456")
	SetCurrencyMapping("ABC", "EFGHIJKL")
}

func TestCurrencyMapping(t *testing.T) {
	if CurrencyMapping("123") != "456" {
		t.Fatal("Error CurrencyMapping")
	} else if CurrencyMapping("ABC") != "EFGHIJKL" {
		t.Fatal("Error CurrencyMapping")
	}
}

func TestCurrencyUnMapping(t *testing.T) {
	if CurrencyUnMapping("456") != "123" {
		t.Fatal("Error CurrencyUnMapping")
	} else if CurrencyUnMapping("EFGHIJKL") != "ABC" {
		t.Fatal("Error CurrencyUnMapping")
	}
}

type testCurrency struct {
	C1 Currency `json:"currency1"`
	C2 Currency `json:"currency2"`
	C3 Currency `json:"currency3"`
	C4 Currency `json:"currency4"`
}

func TestCurrencyJSON(t *testing.T) {
	c := &testCurrency{
		C1: NewCurrency("abc"),
		C2: NewCurrency("ABC"),
		C3: NewCurrency("123"),
		C4: NewCurrency("456"),
	}
	data, err := gojson.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	c = new(testCurrency)
	if err := gojson.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}
	if c.C1.String() != "abc" {
		t.Fatal("Error C1:", c.C1.String())
	} else if c.C2.String() != "EFGHIJKL" {
		t.Fatal("Error C2:", c.C2.String())
	} else if c.C3.String() != "456" {
		t.Fatal("Error C3:", c.C3.String())
	} else if c.C4.String() != "456" {
		t.Fatal("Error C4:", c.C4.String())
	}
}
