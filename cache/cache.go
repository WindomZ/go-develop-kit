package cache

import (
	"github.com/WindomZ/go-develop-kit/cache/freecache"
	"github.com/WindomZ/go-develop-kit/cache/gocache"
	"time"
)

type Cache interface {
	Delete(key string) bool
	SetBytes(key string, value []byte, expireSeconds ...int) error
	GetBytes(key string) ([]byte, error)
	SetString(key string, value string, expireSeconds ...int) error
	GetString(key string) (string, error)
	SetInterface(key string, value interface{}, expireSeconds ...int) error
	GetInterface(key string, values ...interface{}) (interface{}, error)
}

type Config struct {
	MoreString      bool
	MoreInterface   bool
	Size            int
	ExpireSeconds   int
	CleanupInterval int
}

var DefaultConfig *Config = &Config{
	MoreString:      true,
	MoreInterface:   false,
	Size:            1024,
	ExpireSeconds:   60,
	CleanupInterval: 30,
}

func NewFreeCache(size, expireSeconds int) *freecache.Cache {
	return freecache.NewCache(size, expireSeconds)
}

func NewGoCache(expireSeconds, cleanupSeconds int) *gocache.Cache {
	return gocache.NewCache(time.Duration(expireSeconds)*time.Second,
		time.Duration(cleanupSeconds)*time.Second)
}

func NewCache(c *Config) Cache {
	if c == nil {
		return NewCache(DefaultConfig)
	} else if c.MoreString {
		return NewFreeCache(c.Size, c.ExpireSeconds)
	} else if c.MoreInterface {
		return NewGoCache(c.ExpireSeconds, c.CleanupInterval)
	}
	return NewFreeCache(c.Size, c.ExpireSeconds)
}
