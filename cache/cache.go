package cache

type ICache interface {
	SetString(key string, value string, expireSeconds int) error
	GetString(key string) (string, error)
	SetInterface(key string, value interface{}, expireSeconds int) error
	GetInterface(key string, value interface{}) error
}
