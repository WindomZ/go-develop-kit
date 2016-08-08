package boolcache

import (
	"runtime"
	"sync"
	"time"
)

const (
	NoExpiration time.Duration = -1
)

type Item struct {
	Bool       bool
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
	*cache
}

type cache struct {
	defaultExpiration time.Duration
	mux               sync.RWMutex
	items             map[string]Item
	janitor           *janitor
}

func (c *cache) SetBool(k string, v bool, ds ...time.Duration) {
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
		Bool:       v,
		Expiration: e,
	}
	c.mux.Unlock()
}

func (c *cache) GetBool(k string) (bool, bool) {
	c.mux.RLock()
	item, ok := c.items[k]
	c.mux.RUnlock()
	if ok && !item.Expired() {
		return item.Bool, true
	}
	return false, false
}

func (c *cache) SwapBool(k string, ds ...time.Duration) bool {
	var e int64
	d := c.defaultExpiration
	if ds != nil && len(ds) != 0 {
		d = ds[0]
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mux.RLock()
	item, ok := c.items[k]
	if ok {
		item.Bool = !item.Bool
		if e > 0 {
			item.Expiration = e
		}
	} else {
		item = Item{
			Bool:       true,
			Expiration: e,
		}
	}
	c.items[k] = item
	c.mux.RUnlock()
	return item.Bool
}

func (c *cache) UpdateBool(k string, d time.Duration) bool {
	if d <= 0 {
		c.Delete(k)
		return false
	}
	c.mux.RLock()
	item, ok := c.items[k]
	if ok {
		item.Expiration = time.Now().Add(d).UnixNano()
	}
	c.items[k] = item
	c.mux.RUnlock()
	return ok
}

func (c *cache) DeleteExpired() {
	now := time.Now().UnixNano()
	c.mux.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			c.delete(k)
		}
	}
	c.mux.Unlock()
}

func (c *cache) Delete(k string) (bool, bool) {
	c.mux.Lock()
	v, evicted := c.delete(k)
	c.mux.Unlock()
	return v, evicted
}

func (c *cache) delete(k string) (bool, bool) {
	if v, found := c.items[k]; found {
		delete(c.items, k)
		return v.Bool, true
	}
	return false, false
}

type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func (j *janitor) Run(c *cache) {
	j.stop = make(chan bool)
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

func stopJanitor(c *Cache) {
	c.janitor.stop <- true
	c.janitor = nil
}

func runJanitor(c *cache, ci time.Duration) {
	c.janitor = &janitor{
		Interval: ci,
	}
	go c.janitor.Run(c)
}

func newCache(defaultExpiration time.Duration) *cache {
	if defaultExpiration <= 0 {
		defaultExpiration = NoExpiration
	}
	c := &cache{
		defaultExpiration: defaultExpiration,
		items:             make(map[string]Item),
	}
	return c
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	c := newCache(defaultExpiration)
	C := &Cache{c}
	if cleanupInterval > 0 {
		runJanitor(c, cleanupInterval)
		runtime.SetFinalizer(C, stopJanitor)
	}
	return C
}
