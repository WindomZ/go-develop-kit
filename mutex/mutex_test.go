package mutex

import "testing"

var mux *Mutex = NewMutex()

func TestMutexPlan1(t *testing.T) {
	mux.MustLock()
	for i := 0; i < 110; i++ {
		mux.UnsafeLock()
	}
	for i := 0; i < 100; i++ {
		mux.UnsafeUnlock()
	}
	mux.MustUnlock()
}

func TestMutexPlan2(t *testing.T) {
	mux.MustLock()
	for i := 0; i < 100; i++ {
		mux.UnsafeLock()
	}
	for i := 0; i < 110; i++ {
		mux.UnsafeUnlock()
	}
	mux.SafeUnlock()
}

func TestMutexPlan3(t *testing.T) {
	for i := 0; i < 100; i++ {
		mux.UnsafeLock()
		mux.UnsafeUnlock()
		mux.MustLock()
		mux.MustUnlock()
	}
	for i := 0; i < 100; i++ {
		mux.MustLock()
		mux.UnsafeUnlock()
		mux.UnsafeLock()
		mux.MustUnlock()
	}
	for i := 0; i < 100; i++ {
		mux.MustLock()
		mux.UnsafeLock()
		mux.UnsafeLock()
		mux.MustUnlock()
	}
}
