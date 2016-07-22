package cache

import (
	"testing"
	"time"
)

var c *Cache

func TestNewCache(t *testing.T) {
	c = NewCache(time.Minute)
}

func TestCache_Int64(t *testing.T) {
	const KEY string = "TestInt"
	const VALUE int64 = 100
	c.SetInt64(KEY, VALUE)
	if v, ok := c.GetInt64(KEY); !ok {
		t.Fatal("Error key")
	} else if v != VALUE {
		t.Fatal("Error value")
	}
}

func TestCache_Float64(t *testing.T) {
	const KEY string = "TestInt"
	const VALUE float64 = 100.234
	c.SetFloat64(KEY, VALUE)
	if v, ok := c.GetFloat64(KEY); !ok {
		t.Fatal("Error key")
	} else if v != VALUE {
		t.Fatal("Error value")
	}
}
