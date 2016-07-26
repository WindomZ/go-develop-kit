package mutex

import (
	"sync"
	"sync/atomic"
)

type Mutex struct {
	sync.Mutex
	count int32
}

func NewMutex() *Mutex {
	return new(Mutex)
}

// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *Mutex) Lock() {
	m.MustLock()
}

// Same as Lock()
func (m *Mutex) MustLock() {
	m.Mutex.Lock()
	atomic.StoreInt32(&m.count, 1)
}

func (m *Mutex) SafeLock() {
	if atomic.SwapInt32(&m.count, 1) <= 0 {
		m.Mutex.Lock()
	}
}

func (m *Mutex) UnsafeLock() {
	if atomic.CompareAndSwapInt32(&m.count, 0, 1) {
		m.Mutex.Lock()
	} else if atomic.LoadInt32(&m.count) >= 1 {
		atomic.AddInt32(&m.count, 1)
	}
}

// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked Mutex is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func (m *Mutex) Unlock() {
	m.MustUnlock()
}

// Same as Unlock()
func (m *Mutex) MustUnlock() {
	m.Mutex.Unlock()
	atomic.StoreInt32(&m.count, 0)
}

func (m *Mutex) SafeUnlock() {
	if atomic.SwapInt32(&m.count, 0) >= 1 {
		m.Mutex.Unlock()
	}
}

func (m *Mutex) UnsafeUnlock() {
	if atomic.CompareAndSwapInt32(&m.count, 1, 0) {
		m.Mutex.Unlock()
	} else if atomic.LoadInt32(&m.count) > 1 {
		atomic.AddInt32(&m.count, -1)
	}
}
