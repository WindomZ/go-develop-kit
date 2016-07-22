package cache

import "github.com/WindomZ/go-develop-kit/cache/freecache"

type ICache interface {
	SetBytes(key string, value []byte, expireSeconds ...int) error
	GetBytes(key string) ([]byte, error)
	SetString(key string, value string, expireSeconds ...int) error
	GetString(key string) (string, error)
	SetInterface(key string, value interface{}, expireSeconds ...int) error
	GetInterface(key string, value interface{}) (interface{}, error)
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
	CleanupInterval: 5,
}

func NewFreeCache(size, expireSeconds int) *freecache.Cache {
	return freecache.NewCache(size, expireSeconds)
}

func NewCache(c *Config) ICache {
	if c == nil {
		return NewCache(DefaultConfig)
	}
	if c.MoreString {
		return NewFreeCache(c.Size, c.ExpireSeconds)
	} else if c.MoreInterface {
		//TODO: gocache
	}
	return NewFreeCache(1024, c.ExpireSeconds)
}
