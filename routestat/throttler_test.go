package routestat

import (
	"testing"
	"time"
)

var throttler *Throttler

func TestNewThrottler(t *testing.T) {
	throttler = NewThrottler(100 * time.Millisecond)
	//t.Log(throttler.defaultRate)
}

func TestThrottler_Accept(t *testing.T) {
	const IP string = "test"
	var cnt int64 = 0
	for throttler.Accept(IP) {
		cnt++
	}
	time.Sleep(100 * time.Millisecond)
	for throttler.Accept(IP) {
		cnt++
	}
	time.Sleep(199 * time.Millisecond)
	for throttler.Accept(IP) {
		cnt++
	}
	t.Log(cnt)
}

func TestThrottler_Close(t *testing.T) {
	if err := throttler.Close(); err != nil {
		t.Fatal(err)
	}
}
