package unit

import (
	"github.com/WindomZ/go-macro/json"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	SetCurrencyMapping("$", "USD")
	SetCurrencyMapping("RMB", "CNY")
}

func TestSetCurrencyMappingFunc(t *testing.T) {
	SetCurrencyMappingFunc(func(s string) string {
		switch s {
		case "AAA":
			return "bbb"
		}
		return s
	})
	SetCurrencyUnMappingFunc(func(s string) string {
		switch s {
		case "bbb":
			return "AAA"
		}
		return s
	})
	if CurrencyMapping("AAA") != "bbb" {
		t.Fatal("Error SetCurrencyMappingFunc")
	} else if CurrencyUnMapping("bbb") != "AAA" {
		t.Fatal("Error SetCurrencyUnMappingFunc")
	}
}

func TestCurrencyMapping(t *testing.T) {
	if CurrencyMapping("$") != "USD" {
		t.Fatal("Error CurrencyMapping")
	} else if CurrencyMapping("RMB") != "CNY" {
		t.Fatal("Error CurrencyMapping")
	}
}

func TestCurrencyUnMapping(t *testing.T) {
	if CurrencyUnMapping("USD") != "$" {
		t.Fatal("Error CurrencyUnMapping")
	} else if CurrencyUnMapping("CNY") != "RMB" {
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
		C1: NewCurrency("rmb"),
		C2: NewCurrency("RMB"),
		C3: NewCurrency("$"),
		C4: NewCurrency("USD"),
	}
	data, err := gojson.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	c = new(testCurrency)
	if err := gojson.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}
	if c.C1.String() != "rmb" {
		t.Fatal("Error C1:", c.C1.String())
	} else if c.C2.String() != "CNY" {
		t.Fatal("Error C2:", c.C2.String())
	} else if c.C3.String() != "USD" {
		t.Fatal("Error C3:", c.C3.String())
	} else if c.C4.String() != "USD" {
		t.Fatal("Error C4:", c.C4.String())
	}
	var m map[string]string
	if err := gojson.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	for _, v := range m {
		switch v {
		case "rmb":
		case "RMB":
		case "$":
		default:
			t.Fatal("Error Map:", v)
		}
	}
}
