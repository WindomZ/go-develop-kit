package numcache

import (
	"runtime"
	"sync"
	"time"
)

const (
	NoExpiration time.Duration = -1
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
	*cache
}

type cache struct {
	defaultExpiration time.Duration
	mux               sync.RWMutex
	items             map[string]Item
	janitor           *janitor
}

func (c *cache) DefaultExpiration() time.Duration {
	return c.defaultExpiration
}

func (c *cache) SetInt64(k string, v int64, ds ...time.Duration) {
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

func (c *cache) GetInt64(k string) (int64, bool) {
	c.mux.RLock()
	item, ok := c.items[k]
	c.mux.RUnlock()
	if ok && !item.Expired() {
		return item.Int, true
	}
	return 0, false
}

func (c *cache) SetFloat64(k string, v float64, ds ...time.Duration) {
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

func (c *cache) GetFloat64(k string) (float64, bool) {
	c.mux.RLock()
	item, ok := c.items[k]
	c.mux.RUnlock()
	if ok && !item.Expired() {
		return item.Float, true
	}
	return 0, false
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

func (c *cache) Delete(k string) bool {
	c.mux.Lock()
	evicted := c.delete(k)
	c.mux.Unlock()
	return evicted
}

func (c *cache) delete(k string) bool {
	if _, found := c.items[k]; found {
		delete(c.items, k)
		return true
	}
	return false
}

func (c *cache) Update(k string, d time.Duration) bool {
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

func (c *cache) IncrementInt64(k string, v int64, ds ...time.Duration) int64 {
	var e int64
	d := c.defaultExpiration
	if ds != nil && len(ds) != 0 {
		d = ds[0]
	}
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

func (c *cache) DecrementInt64(k string, v int64, ds ...time.Duration) int64 {
	return c.IncrementInt64(k, -v, ds...)
}

func (c *cache) IncrementFloat64(k string, v float64, ds ...time.Duration) float64 {
	var e int64
	d := c.defaultExpiration
	if ds != nil && len(ds) != 0 {
		d = ds[0]
	}
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

func (c *cache) DecrementFloat64(k string, v float64, ds ...time.Duration) float64 {
	return c.IncrementFloat64(k, -v, ds...)
}

func (c *cache) MapInt64() map[string]int64 {
	c.mux.Lock()
	r := make(map[string]int64, len(c.items))
	for k, v := range c.items {
		r[k] = v.Int
	}
	c.mux.Unlock()
	return r
}

func (c *cache) MapFloat64() map[string]float64 {
	c.mux.Lock()
	r := make(map[string]float64, len(c.items))
	for k, v := range c.items {
		r[k] = v.Float
	}
	c.mux.Unlock()
	return r
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

func NewCacheSameExpiration(defaultExpiration time.Duration) *Cache {
	return NewCache(defaultExpiration, NoExpiration)
}

func NewCacheNoExpiration() *Cache {
	return NewCache(NoExpiration, NoExpiration)
}
