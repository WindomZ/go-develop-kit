package numcache

import (
	"sync"
	"time"
)

type Item struct {
	Int        int64
	Float      float64
	Expiration int64
}

// Returns true if the item has expired.
func (item *Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	defaultExpiration time.Duration
	mux               sync.RWMutex
	items             map[string]Item
}

func NewCache(defaultExpiration time.Duration) *Cache {
	return &Cache{
		defaultExpiration: defaultExpiration,
		items:             make(map[string]Item),
	}
}

func (c *Cache) SetInt64(k string, v int64, ds ...time.Duration) {
	var e int64
	d := c.defaultExpiration
	if ds != nil && len(ds) != 0 {
		d = ds[0]
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mux.Lock()
	c.items[k] = Item{
		Int:        v,
		Expiration: e,
	}
	c.mux.Unlock()
}

func (c *Cache) GetInt64(k string) (int64, bool) {
	c.mux.RLock()
	item, ok := c.items[k]
	c.mux.RUnlock()
	if ok && !item.Expired() {
		return item.Int, true
	}
	return 0, false
}

func (c *Cache) SetFloat64(k string, v float64, ds ...time.Duration) {
	var e int64
	d := c.defaultExpiration
	if ds != nil && len(ds) != 0 {
		d = ds[0]
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mux.Lock()
	c.items[k] = Item{
		Float:      v,
		Expiration: e,
	}
	c.mux.Unlock()
}

func (c *Cache) GetFloat64(k string) (float64, bool) {
	c.mux.RLock()
	item, ok := c.items[k]
	c.mux.RUnlock()
	if ok && !item.Expired() {
		return item.Float, true
	}
	return 0, false
}

func (c *Cache) IncrementInt64(k string, v int64, d time.Duration) int64 {
	var e int64
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mux.Lock()
	item, ok := c.items[k]
	if ok && !item.Expired() {
		item.Int += v
		if e != 0 {
			item.Expiration = e
		}
	} else {
		item = Item{
			Int:        v,
			Expiration: e,
		}
	}
	c.items[k] = item
	c.mux.Unlock()
	return item.Int
}

func (c *Cache) DecrementInt64(k string, v int64, d time.Duration) int64 {
	return c.IncrementInt64(k, -v, d)
}

func (c *Cache) IncrementFloat64(k string, v float64, d time.Duration) float64 {
	var e int64
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mux.Lock()
	item, ok := c.items[k]
	if ok && !item.Expired() {
		item.Float += v
		if e != 0 {
			item.Expiration = e
		}
	} else {
		item = Item{
			Float:      v,
			Expiration: e,
		}
	}
	c.items[k] = item
	c.mux.Unlock()
	return item.Float
}

func (c *Cache) DecrementFloat64(k string, v float64, d time.Duration) float64 {
	return c.IncrementFloat64(k, -v, d)
}
