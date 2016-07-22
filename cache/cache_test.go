package cache

import (
	"testing"
	"time"
)

var c Cache

func TestNewCache(t *testing.T) {
	c = NewCache(&Config{
		MoreString:      false,
		MoreInterface:   true,
		Size:            1024,
		ExpireSeconds:   60,
		CleanupInterval: 30,
	})
	if c == nil {
		t.Fatal("Error create")
	}
}

func TestCache_String(t *testing.T) {
	var key, value string = "key", "test 123 ABC 中文 !@#"
	if err := c.SetString(key, value); err != nil {
		t.Fatal(err)
	} else if v, err := c.GetString(key); err != nil {
		t.Fatal(err)
	} else if v != value {
		t.Fatal("Diffence string")
	}
}

type TestDemo1 struct {
	Int    int       `json:"int"`
	String string    `json:"string"`
	Time   time.Time `json:"time"`
}

func TestCache_Interface(t *testing.T) {
	var key, value string = "key", "test 123 ABC 中文 !@#"
	d1 := &TestDemo1{
		Int:    1,
		String: value,
		Time:   time.Now(),
	}
	if err := c.SetInterface(key, d1); err != nil {
		t.Fatal(err)
	}
	var d2 TestDemo1
	if _, err := c.GetInterface(key, &d2); err != nil {
		t.Fatal(err)
	} else if d1.Int != d2.Int {
		t.Fatal("Diffence int")
	} else if d1.String != d2.String {
		t.Fatal("Diffence string")
	} else if !d1.Time.Equal(d2.Time) {
		t.Fatal("Diffence time")
	}
}
