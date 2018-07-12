package uuid

import (
	"strconv"
	"sync"

	"github.com/google/uuid"
)

var (
	idx uint64 = 0
	mux        = new(sync.Mutex)
)

// NewUUID returns random generated UUID.
func NewUUID() string {
	if id, err := uuid.NewRandom(); err == nil {
		return id.String()
	}
	return ""
}

// NewUUIDWithName returns UUID based on SHA-1 hash of namespace UUID and name.
func NewUUIDWithName(name string) string {
	if id, err := uuid.NewRandom(); err == nil {
		return uuid.NewSHA1(id, []byte(name)).String()
	}
	return ""
}

// NewSafeUUID returns random generated UUID safely.
func NewSafeUUID() (r string) {
	mux.Lock()
	idx++
	if id, err := uuid.NewRandom(); err == nil {
		r = uuid.NewSHA1(id, []byte(strconv.FormatUint(idx, 10))).String()
	}
	mux.Unlock()
	return
}
