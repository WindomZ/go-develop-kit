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

func (mux *Mutex) MustLock() {
	mux.Lock()
	atomic.StoreInt32(&mux.count, 1)
}

func (mux *Mutex) SafeLock() {
	if atomic.SwapInt32(&mux.count, 1) <= 0 {
		mux.Lock()
	}
}

func (mux *Mutex) UnsafeLock() {
	if atomic.CompareAndSwapInt32(&mux.count, 0, 1) {
		mux.Lock()
	} else if atomic.LoadInt32(&mux.count) >= 1 {
		atomic.AddInt32(&mux.count, 1)
	}
}

func (mux *Mutex) MustUnlock() {
	mux.Unlock()
	atomic.StoreInt32(&mux.count, 0)
}

func (mux *Mutex) SafeUnlock() {
	if atomic.SwapInt32(&mux.count, 0) >= 1 {
		mux.Unlock()
	}
}

func (mux *Mutex) UnsafeUnlock() {
	if atomic.CompareAndSwapInt32(&mux.count, 1, 0) {
		mux.Unlock()
	} else if atomic.LoadInt32(&mux.count) > 1 {
		atomic.AddInt32(&mux.count, -1)
	}
}
