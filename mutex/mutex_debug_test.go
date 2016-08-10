package mutex

import (
	"testing"
	"time"
)

var mux_debug *MutexDebug = NewMutexDebug()

func TestMutexDebug_Plan1(t *testing.T) {
	mux_debug.Lock()
	//t.Log(mux_debug.Log())
	if mux_debug.Line() != 11 {
		t.Fatal(mux_debug.Log())
	}
	mux_debug.Unlock()
}

func TestMutexDebug_Plan2(t *testing.T) {
	mux_debug.Lock()
	time.Sleep(time.Second)
	if mux_debug.DurationSeconds() != 1 {
		t.Fatal("Error plan2 lock", mux_debug.Duration())
	}
	mux_debug.Unlock()
	if mux_debug.Duration() != 0 {
		t.Fatal("Error plan2 unlock", mux_debug.Duration())
	}
}
