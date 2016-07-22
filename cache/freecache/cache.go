package freecache

import (
	"encoding/json"
	"github.com/coocood/freecache"
)

type Cache struct {
	base          freecache.Cache
	ExpireSeconds int
}

func NewCache(size, expireSeconds int) *Cache {
	return &Cache{
		base:          *freecache.NewCache(size),
		ExpireSeconds: expireSeconds,
	}
}

func (c *Cache) Ex() *freecache.Cache {
	return &c.base
}

func (c *Cache) SetBytes(key string, value []byte, expireSeconds ...int) error {
	if len(key) == 0 {
		return ErrNoKey
	}
	if expireSeconds != nil && len(expireSeconds) != 0 {
		return c.Ex().Set([]byte(key), value, expireSeconds[0])
	}
	return c.Ex().Set([]byte(key), value, c.ExpireSeconds)
}

func (c *Cache) GetBytes(key string) ([]byte, error) {
	if len(key) == 0 {
		return []byte{}, ErrNoKey
	}
	return c.Ex().Get([]byte(key))
}

func (c *Cache) SetString(key string, value string, expireSeconds ...int) error {
	return c.SetBytes(key, []byte(value), expireSeconds...)
}

func (c *Cache) GetString(key string) (string, error) {
	if v, err := c.GetBytes(key); err != nil {
		return "", decorateError(err)
	} else {
		return string(v), nil
	}
}

func (c *Cache) SetInterface(key string, value interface{}, expireSeconds ...int) error {
	if value == nil {
		return ErrNoValue
	}
	v, err := json.Marshal(value)
	if err != nil {
		return decorateError(err)
	}
	err = c.SetBytes(key, v, expireSeconds...)
	return nil
}

func (c *Cache) GetInterface(key string, value interface{}) error {
	if value == nil {
		return ErrNoValue
	}
	v, err := c.GetBytes(key)
	if err != nil {
		return decorateError(err)
	}
	return json.Unmarshal(v, value)
}
