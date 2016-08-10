package mutex

import "testing"

var mux *Mutex = NewMutex()

func TestMutex_Plan1(t *testing.T) {
	for i := 0; i < 100; i++ {
		mux.Lock()
		mux.Unlock()
	}
}

func TestMutex_Plan2(t *testing.T) {
	if mux.IsLocked() {
		t.Fatal("Error plan1")
	}
	mux.Lock()
	if !mux.IsLocked() {
		t.Fatal("Error plan2 lock")
	}
	mux.Unlock()
	if mux.IsLocked() {
		t.Fatal("Error plan2 unlock")
	}
}
