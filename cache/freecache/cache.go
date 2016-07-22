package freecache

import (
	"encoding/json"
	"github.com/coocood/freecache"
)

type Cache struct {
	freecache.Cache
}

func NewCache(size int) *Cache {
	return &Cache{
		Cache: *freecache.NewCache(size),
	}
}

func (c *Cache) SetString(key string, value string, expireSeconds int) error {
	if len(key) == 0 {
		return ErrNoKey
	}
	return c.Set([]byte(key), []byte(value), expireSeconds)
}

func (c *Cache) GetString(key string) (string, error) {
	if len(key) == 0 {
		return "", ErrNoKey
	} else if v, err := c.Get([]byte(key)); err != nil {
		return "", decorateError(err)
	} else {
		return string(v), nil
	}
}

func (c *Cache) SetInterface(key string, value interface{}, expireSeconds int) error {
	if len(key) == 0 {
		return ErrNoKey
	} else if value == nil {
		return ErrNoValue
	}
	v, err := json.Marshal(value)
	if err != nil {
		return decorateError(err)
	}
	err = c.Set([]byte(key), v, expireSeconds)
	return nil
}

func (c *Cache) GetInterface(key string, value interface{}) error {
	if len(key) == 0 {
		return ErrNoKey
	} else if value == nil {
		return ErrNoValue
	}
	v, err := c.Get([]byte(key))
	if err != nil {
		return decorateError(err)
	}
	return json.Unmarshal(v, value)
}
