package gocache

import (
	"github.com/WindomZ/go-struct-filler"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	base          cache.Cache
	ExpireSeconds time.Duration
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	return &Cache{
		base:          *cache.New(defaultExpiration, cleanupInterval),
		ExpireSeconds: defaultExpiration,
	}
}

func (c *Cache) Ex() *cache.Cache {
	return &c.base
}

func (c *Cache) SetBytes(key string, value []byte, expireSeconds ...int) error {
	return c.SetInterface(key, value, expireSeconds...)
}

func (c *Cache) GetBytes(key string) ([]byte, error) {
	if v, err := c.GetInterface(key, ""); err != nil {
		return []byte{}, err
	} else if data, ok := v.([]byte); ok {
		return data, nil
	}
	return []byte{}, ErrNotExistKey
}

func (c *Cache) SetString(key string, value string, expireSeconds ...int) error {
	return c.SetInterface(key, value, expireSeconds...)
}

func (c *Cache) GetString(key string) (string, error) {
	if v, err := c.GetInterface(key, ""); err != nil {
		return "", err
	} else if str, ok := v.(string); ok {
		return str, nil
	}
	return "", ErrNotExistKey
}

func (c *Cache) SetInterface(key string, value interface{}, expireSeconds ...int) error {
	if len(key) == 0 {
		return ErrNoKey
	} else if value == nil {
		return ErrNoValue
	}
	if expireSeconds != nil && len(expireSeconds) != 0 {
		c.Ex().Set(key, value, time.Duration(expireSeconds[0])*time.Second)
	} else {
		c.Ex().Set(key, value, c.ExpireSeconds)
	}
	return nil
}

func (c *Cache) GetInterface(key string, value interface{}) (interface{}, error) {
	if len(key) == 0 {
		return value, ErrNoKey
	} else if v, ok := c.Ex().Get(key); ok {
		if value != nil {
			gsf.ConvertStruct(v, value)
		}
		return v, nil
	}
	return value, ErrNotExistKey
}
