package mutex

import (
	"sync"
	"sync/atomic"
)

type Mutex struct {
	sync.Mutex
	before int32
	after  int32
}

func NewMutex() *Mutex {
	return new(Mutex)
}

// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *Mutex) lock() {
	atomic.StoreInt32(&m.before, 1)
	m.Mutex.Lock()
	atomic.StoreInt32(&m.after, 1)
}

// Same as lock()
func (m *Mutex) Lock() {
	m.lock()
}

// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked Mutex is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func (m *Mutex) unlock() {
	atomic.StoreInt32(&m.after, 0)
	m.Mutex.Unlock()
	atomic.StoreInt32(&m.before, 0)
}

// Same as unlock()
func (m *Mutex) Unlock() {
	m.unlock()
}

func (m *Mutex) IsLocked() bool {
	return atomic.LoadInt32(&m.before) == 1 &&
		atomic.LoadInt32(&m.after) == 1
}
