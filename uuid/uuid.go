package uuid

import (
	"github.com/satori/go.uuid"
	"sync"
)

var (
	idx uint64      = 0
	mux *sync.Mutex = new(sync.Mutex)
)

var EMPTY, _ = uuid.FromString("00000000-0000-0000-0000-000000000000")

func IsEmpty(id string) bool {
	u, err := uuid.FromString(id)
	if err != nil {
		return false
	}
	return uuid.Equal(EMPTY, u)
}

func NewUUID() string {
	return uuid.NewV4().String()
}

func NewUUIDWithName(name string) string {
	return uuid.NewV5(uuid.NewV4(), name).String()
}

func NewSafeUUID() string {
	mux.Lock()
	idx++
	r := uuid.NewV5(uuid.NewV4(), string(idx)).String()
	mux.Unlock()
	return r
}
