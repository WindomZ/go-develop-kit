package mutex

import "sync/atomic"

type MutexSafe struct {
	Mutex
	count int32
}

func NewMutexSafe() *MutexSafe {
	return new(MutexSafe)
}

// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *MutexSafe) lock() {
	m.Mutex.Lock()
}

func (m *MutexSafe) Lock() {
	m.MustLock()
}

func (m *MutexSafe) MustLock() {
	m.lock()
	atomic.StoreInt32(&m.count, 1)
}

func (m *MutexSafe) SafeLock() {
	if atomic.SwapInt32(&m.count, 1) <= 0 {
		m.lock()
	}
}

func (m *MutexSafe) UnsafeLock() {
	if atomic.CompareAndSwapInt32(&m.count, 0, 1) {
		m.lock()
	} else if atomic.LoadInt32(&m.count) >= 1 {
		atomic.AddInt32(&m.count, 1)
	}
}

// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked MutexSafe is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a MutexSafe and then
// arrange for another goroutine to unlock it.
func (m *MutexSafe) unlock() {
	m.Mutex.Unlock()
}

func (m *MutexSafe) Unlock() {
	m.MustUnlock()
}

func (m *MutexSafe) MustUnlock() {
	m.unlock()
	atomic.StoreInt32(&m.count, 0)
}

func (m *MutexSafe) SafeUnlock() {
	if atomic.SwapInt32(&m.count, 0) >= 1 {
		m.unlock()
	}
}

func (m *MutexSafe) UnsafeUnlock() {
	if atomic.CompareAndSwapInt32(&m.count, 1, 0) {
		m.unlock()
	} else if atomic.LoadInt32(&m.count) > 1 {
		atomic.AddInt32(&m.count, -1)
	}
}
