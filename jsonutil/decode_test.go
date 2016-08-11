package jsonutil

import (
	"encoding/json"
	"testing"
)

type testLayer1 struct {
	testLayer2   `json:""`
	String1      string     `json:"string1"`
	Int1         int64      `json:"int1"`
	Float1       float64    `json:"float1"`
	Bool1        bool       `json:"bool1"`
	ArrayString1 []string   `json:"strings1"`
	Layer3       testLayer3 `json:""`
}

type testLayer2 struct {
	String2      string   `json:"string2"`
	Int2         int64    `json:"int2"`
	Float2       float64  `json:"float2"`
	Bool2        bool     `json:"bool2"`
	ArrayString2 []string `json:"strings2"`
}

type testLayer3 struct {
	String3      string   `json:"string3"`
	Int3         int64    `json:"int3"`
	Float3       float64  `json:"float3"`
	Bool3        bool     `json:"bool3"`
	ArrayString3 []string `json:"strings3"`
}

var testLayer *testLayer1

func TestJSONMustInit(t *testing.T) {
	testLayer = &testLayer1{
		testLayer2: testLayer2{
			String2: "S2",
			Int2:    2,
			Float2:  2.2,
			Bool2:   true,
			ArrayString2: []string{
				"Ss21",
				"Ss22",
				"Ss23",
			},
		},
		String1: "S1",
		Int1:    1,
		Float1:  1.1,
		Bool1:   false,
		ArrayString1: []string{
			"Ss11",
			"Ss12",
			"Ss13",
		},
		Layer3: testLayer3{
			String3: "S3",
			Int3:    3,
			Float3:  3.3,
			Bool3:   true,
			ArrayString3: []string{
				"Ss31",
				"Ss32",
				"Ss33",
			},
		},
	}
}

func TestJSONMustGetString(t *testing.T) {
	data, err := json.Marshal(testLayer)
	if err != nil {
		t.Fatal(err)
	}
	if s := JSONMustGetString(data, "string1"); s != "S1" {
		t.Fatal("Error string1", s)
	}
	if s := JSONMustGetString(data, "string2"); s != "S2" {
		t.Fatal("Error string2", s)
	}
	if s := JSONMustGetString(data, "Layer3", "string3"); s != "S3" {
		t.Fatal("Error string3", s)
	}
}

func TestJSONMustGetInt(t *testing.T) {
	data, err := json.Marshal(testLayer)
	if err != nil {
		t.Fatal(err)
	}
	if i := JSONMustGetInt(data, "int1"); i != 1 {
		t.Fatal("Error int1", i)
	}
	if i := JSONMustGetInt(data, "int2"); i != 2 {
		t.Fatal("Error int2", i)
	}
	if i := JSONMustGetInt(data, "Layer3", "int3"); i != 3 {
		t.Fatal("Error int3", i)
	}
}

func TestJSONMustGetFloat(t *testing.T) {
	data, err := json.Marshal(testLayer)
	if err != nil {
		t.Fatal(err)
	}
	if i := JSONMustGetFloat(data, "float1"); i != 1.1 {
		t.Fatal("Error float1", i)
	}
	if i := JSONMustGetFloat(data, "float2"); i != 2.2 {
		t.Fatal("Error float2", i)
	}
	if i := JSONMustGetFloat(data, "Layer3", "float3"); i != 3.3 {
		t.Fatal("Error float3", i)
	}
}

func TestJSONMustGetBoolean(t *testing.T) {
	data, err := json.Marshal(testLayer)
	if err != nil {
		t.Fatal(err)
	}
	if b := JSONMustGetBoolean(data, "bool1"); b {
		t.Fatal("Error bool1", b)
	}
	if b := JSONMustGetBoolean(data, "bool2"); !b {
		t.Fatal("Error bool2", b)
	}
	if b := JSONMustGetBoolean(data, "bool3"); b {
		t.Fatal("Error bool3", b)
	}
}
