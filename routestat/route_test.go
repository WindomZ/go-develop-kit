package routestat

import (
	"testing"
	"time"
)

var routeStat *RouteStat

func TestNewRouteStat(t *testing.T) {
	routeStat = NewRouteStat(time.Second, 100*time.Millisecond)
}

func TestRouteStat_Start(t *testing.T) {
	if err := routeStat.Start(); err != nil {
		t.Fatal(err)
	}
}

func TestRouteStat_Accept(t *testing.T) {
	const IP string = "test"
	var cnt int64 = 0
	for routeStat.Accept(IP) {
		cnt++
	}
	time.Sleep(100 * time.Millisecond)
	for routeStat.Accept(IP) {
		cnt++
	}
	time.Sleep(199 * time.Millisecond)
	for routeStat.Accept(IP) {
		cnt++
	}
	t.Log(cnt)
}

func TestRouteStat_String(t *testing.T) {
	time.Sleep(time.Second)
	t.Log(routeStat.String())
}

func TestRouteStat_Close(t *testing.T) {
	if err := routeStat.Close(); err != nil {
		t.Fatal(err)
	}
}
