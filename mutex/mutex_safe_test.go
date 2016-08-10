package mutex

import "testing"

var mux_safe *MutexSafe = NewMutexSafe()

func TestMutexSafe_Plan1(t *testing.T) {
	mux_safe.MustLock()
	for i := 0; i < 110; i++ {
		mux_safe.UnsafeLock()
	}
	for i := 0; i < 100; i++ {
		mux_safe.UnsafeUnlock()
	}
	mux_safe.MustUnlock()
}

func TestMutexSafe_Plan2(t *testing.T) {
	if mux_safe.IsLocked() {
		t.Fatal("Error plan1")
	}
	mux_safe.MustLock()
	for i := 0; i < 100; i++ {
		mux_safe.UnsafeLock()
	}
	for i := 0; i < 110; i++ {
		mux_safe.UnsafeUnlock()
	}
	mux_safe.SafeUnlock()
}

func TestMutexSafe_Plan3(t *testing.T) {
	if mux_safe.IsLocked() {
		t.Fatal("Error plan2")
	}
	for i := 0; i < 100; i++ {
		mux_safe.UnsafeLock()
		mux_safe.UnsafeUnlock()
		mux_safe.MustLock()
		mux_safe.MustUnlock()
	}
	for i := 0; i < 100; i++ {
		mux_safe.MustLock()
		mux_safe.UnsafeUnlock()
		mux_safe.UnsafeLock()
		mux_safe.MustUnlock()
	}
	for i := 0; i < 100; i++ {
		mux_safe.MustLock()
		mux_safe.UnsafeLock()
		mux_safe.UnsafeLock()
		mux_safe.MustUnlock()
	}
}
