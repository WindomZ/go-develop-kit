package boolcache

import (
	"testing"
	"time"
)

var c *Cache

const (
	_KEY1 string = "Test"
	_KEY2        = ""
)

func TestNewCache(t *testing.T) {
	c = NewCache(time.Minute, time.Second)
}

func TestCache_SetBool(t *testing.T) {
	c.SetBool(_KEY1, true)
	c.SetBool(_KEY2, true)
}

func TestCache_GetBool(t *testing.T) {
	if b, ok := c.GetBool(_KEY1); !b || !ok {
		t.Fatal("Error SetBool")
	}
	if b, ok := c.GetBool(_KEY2); !b || !ok {
		t.Fatal("Error SetBool")
	}
}

func TestCache_SwapBool(t *testing.T) {
	if b := c.SwapBool(_KEY1); b {
		t.Fatal("Error SwapBool")
	}
	if b := c.SwapBool(_KEY2); b {
		t.Fatal("Error SwapBool")
	}
	if b, ok := c.GetBool(_KEY1); b || !ok {
		t.Fatal("Error SwapBool")
	}
	if b, ok := c.GetBool(_KEY2); b || !ok {
		t.Fatal("Error SwapBool")
	}
}

func TestCache_UpdateBool(t *testing.T) {
	if ok := c.UpdateBool(_KEY1, time.Second); !ok {
		t.Fatal("Error UpdateBool")
	}
	if ok := c.UpdateBool(_KEY2, time.Second); !ok {
		t.Fatal("Error UpdateBool")
	}
	time.Sleep(time.Second)
	if _, ok := c.GetBool(_KEY1); ok {
		t.Fatal("Error UpdateBool")
	}
	if _, ok := c.GetBool(_KEY2); ok {
		t.Fatal("Error UpdateBool")
	}
}
