package numcache

import (
	"testing"
	"time"
)

var c *Cache

func TestNewCache(t *testing.T) {
	c = NewCache(time.Minute, time.Second)
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
	const KEY string = "TestFloat"
	const VALUE float64 = 100.234
	c.SetFloat64(KEY, VALUE)
	if v, ok := c.GetFloat64(KEY); !ok {
		t.Fatal("Error key")
	} else if v != VALUE {
		t.Fatal("Error value")
	}
}

func TestCache_Delete(t *testing.T) {
	TestCache_Int64(t)
	const KEY string = "TestInt"
	if _, ok := c.GetInt64(KEY); !ok {
		t.Fatal("Error Delete")
	}
	c.Delete(KEY)
	time.Sleep(time.Second)
	if _, ok := c.GetInt64(KEY); ok {
		t.Fatal("Error Delete")
	}
}

func TestCache_Update(t *testing.T) {
	TestCache_Int64(t)
	const KEY string = "TestInt"
	if _, ok := c.GetInt64(KEY); !ok {
		t.Fatal("Error Update")
	}
	c.Update(KEY, time.Second)
	time.Sleep(time.Second)
	if _, ok := c.GetInt64(KEY); ok {
		t.Fatal("Error Update")
	}
}

func TestCache_IncrementInt64(t *testing.T) {
	const KEY string = "IncrementInt64"
	const VALUE int64 = 1
	for i := int64(0); i < 100; i++ {
		if r := c.IncrementInt64(KEY, VALUE, time.Second); r != (i + 1) {
			t.Fatal("Error IncrementInt64!", r, i)
		}
	}
}

func TestCache_DecrementInt64(t *testing.T) {
	const KEY string = "DecrementInt64"
	const VALUE int64 = 1
	for i := int64(0); i < 100; i++ {
		if r := c.DecrementInt64(KEY, VALUE, time.Second); r != -(i + 1) {
			t.Fatal("Error DecrementInt64!", r, i)
		}
	}
}

func TestCache_IncrementFloat64(t *testing.T) {
	const KEY string = "IncrementFloat64"
	const VALUE float64 = 1
	for i := float64(0); i < 100; i++ {
		if r := c.IncrementFloat64(KEY, VALUE, time.Second); r != (i + 1) {
			t.Fatal("Error IncrementFloat64!", r, i)
		}
	}
}

func TestCache_DecrementFloat64(t *testing.T) {
	const KEY string = "DecrementFloat64"
	const VALUE float64 = 1
	for i := float64(0); i < 100; i++ {
		if r := c.DecrementFloat64(KEY, VALUE, time.Second); r != -(i + 1) {
			t.Fatal("Error DecrementFloat64!", r, i)
		}
	}
}

func TestCache_MapInt64(t *testing.T) {
	const KEY string = "MapInt64"
	const VALUE int64 = 12345
	c.SetInt64(KEY, VALUE)
	m := c.MapInt64()
	if _, ok := m[KEY]; !ok {
		t.Fatal("Error MapInt64!")
	}
}

func TestCache_MapFloat64(t *testing.T) {
	const KEY string = "MapFloat64"
	const VALUE float64 = 1.2345
	c.SetFloat64(KEY, VALUE)
	m := c.MapFloat64()
	if _, ok := m[KEY]; !ok {
		t.Fatal("Error MapFloat64!")
	}
}
