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
	if id, err := uuid.NewV4(); err == nil {
		return id.String()
	}
	return ""
}

// NewUUIDWithName returns UUID based on SHA-1 hash of namespace UUID and name.
func NewUUIDWithName(name string) string {
	if id, err := uuid.NewV4(); err == nil {
		return uuid.NewV5(id, name).String()
	}
	return ""
}

// NewSafeUUID returns random generated UUID safely.
func NewSafeUUID() (r string) {
	mux.Lock()
	idx++
	if id, err := uuid.NewV4(); err == nil {
		r = uuid.NewV5(id, string(idx)).String()
	}
	mux.Unlock()
	return
}
