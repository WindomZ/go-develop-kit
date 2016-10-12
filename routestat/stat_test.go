package routestat

import (
	"testing"
	"time"
)

var stats *Stats

func TestNewStats(t *testing.T) {
	stats = NewStats(time.Second)
}

func TestStats_Start(t *testing.T) {
	if err := stats.Start(); err != nil {
		t.Fatal(err)
	}
}

func TestStats_Accept(t *testing.T) {
	const IP string = "test"
	var cnt int64 = 0
	for stats.Accept(IP) {
		if cnt++; cnt >= 10 {
			break
		}
	}
	t.Log(cnt)
}

func TestStats_String(t *testing.T) {
	time.Sleep(time.Second * 2)
	t.Log(stats.String())
}

func TestStats_Close(t *testing.T) {
	if err := stats.Close(); err != nil {
		t.Fatal(err)
	}
}
