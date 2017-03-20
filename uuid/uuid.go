package uuid

import (
	"sync"

	"github.com/satori/go.uuid"
)

var (
	idx uint64      = 0
	mux *sync.Mutex = new(sync.Mutex)
)

// NewUUID returns random generated UUID.
func NewUUID() string {
	return uuid.NewV4().String()
}

// NewUUIDWithName returns UUID based on SHA-1 hash of namespace UUID and name.
func NewUUIDWithName(name string) string {
	return uuid.NewV5(uuid.NewV4(), name).String()
}

// NewSafeUUID returns random generated UUID safely.
func NewSafeUUID() (r string) {
	mux.Lock()
	idx++
	r = uuid.NewV5(uuid.NewV4(), string(idx)).String()
	mux.Unlock()
	return
}
