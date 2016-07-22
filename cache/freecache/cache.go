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
	var err error
	if expireSeconds != nil && len(expireSeconds) != 0 {
		err = c.Ex().Set([]byte(key), value, expireSeconds[0])
	} else {
		err = c.Ex().Set([]byte(key), value, c.ExpireSeconds)
	}
	if err != nil {
		return decorateError(err)
	}
	return nil
}

func (c *Cache) GetBytes(key string) ([]byte, error) {
	if len(key) == 0 {
		return []byte{}, ErrNoKey
	}
	v, err := c.Ex().Get([]byte(key))
	return v, decorateError(err)
}

func (c *Cache) SetString(key string, value string, expireSeconds ...int) error {
	return c.SetBytes(key, []byte(value), expireSeconds...)
}

func (c *Cache) GetString(key string) (string, error) {
	if v, err := c.GetBytes(key); err != nil {
		return "", err
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

func (c *Cache) GetInterface(key string, value interface{}) (interface{}, error) {
	if value == nil {
		return value, ErrNoValue
	} else if data, err := c.GetBytes(key); err != nil {
		return value, err
	} else if err := json.Unmarshal(data, value); err != nil {
		return value, decorateError(err)
	}
	return value, nil
}
