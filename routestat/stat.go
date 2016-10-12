package routestat

import (
	"encoding/json"
	"github.com/WindomZ/go-develop-kit/cache/freecache"
	"github.com/WindomZ/go-develop-kit/cache/numcache"
	"runtime"
	"strings"
	"sync"
	"time"
)

type IStats interface {
	Start() error
	Record() error
	Accept(ip string) bool
	String() string
	Close() error
}

type Stats struct {
	mutex       sync.RWMutex
	freq        time.Duration
	total       uint64
	timeKey     string
	timeCounter *numcache.Cache
	timeStats   *freecache.Cache
	memStats    runtime.MemStats
	closing     chan struct{}
}

type SimpleMemStats struct {
	Total        uint64 `json:"total"`
	Count        uint64 `json:"count"`
	Mallocs      uint64 `json:"mallocs"`       // number of mallocs
	Frees        uint64 `json:"frees"`         // number of frees
	HeapInuse    uint64 `json:"heap_inuse"`    // bytes in non-idle span
	HeapReleased uint64 `json:"heap_released"` // bytes released to the OS
}

func NewStats(freq time.Duration) *Stats {
	return &Stats{
		freq:        freq,
		timeCounter: numcache.NewCache(time.Hour*24, time.Hour*6),
		timeStats:   freecache.NewCache(512*1024, 86400),
		closing:     make(chan struct{}),
	}
}

func (s *Stats) Start() error {
	s.recordTimeKey()
	go func() {
		ticker := time.NewTicker(s.freq)
		defer ticker.Stop()
		for _ = range ticker.C {
			select {
			case <-s.closing:
				return
			default:
			}
			s.Record()
		}
	}()
	return nil
}

func (s *Stats) recordTimeKey() {
	s.timeKey = time.Now().Format("20060102150405")
}

func (s *Stats) Record() error {
	s.mutex.Lock()
	count, _ := s.timeCounter.GetInt64(s.timeKey)
	s.recordTimeKey()
	var lastMallocs, lastFrees uint64 = s.memStats.Mallocs, s.memStats.Frees
	runtime.ReadMemStats(&s.memStats)
	if data, err := json.Marshal(&SimpleMemStats{
		Total:        s.total,
		Count:        uint64(count),
		Mallocs:      s.memStats.Mallocs - lastMallocs,
		Frees:        s.memStats.Frees - lastFrees,
		HeapInuse:    s.memStats.HeapInuse,
		HeapReleased: s.memStats.HeapReleased,
	}); err == nil {
		s.timeStats.SetString(s.timeKey, strings.Replace(string(data), "\"", "", -1))
	}
	s.mutex.Unlock()
	return nil
}

func (s *Stats) Accept(ip string) bool {
	if len(s.timeKey) != 0 {
		s.timeCounter.IncrementInt64(s.timeKey, 1)
	}
	s.total++
	return true
}

func (s *Stats) String() string {
	str, _ := s.timeStats.GetString(s.timeKey)
	return str
}

func (s *Stats) Close() error {
	close(s.closing)
	return nil
}
