package routestat

import (
	"errors"
	"time"
)

// RouteStat with two modules: Stats and Throttler.
// Must pass the function NewRouteStat to create
type RouteStat struct {
	stats     *Stats
	throttler *Throttler
}

// NewRouteStat returns a RouteStat with a single filler go-routine for all
// throttler which ticks every tbFreq, and record stats which ticks every statFreq.
func NewRouteStat(statFreq, tbFreq time.Duration) *RouteStat {
	return &RouteStat{
		stats:     NewStats(statFreq),
		throttler: NewThrottler(tbFreq),
	}
}

// Start start service, if want Stats module to take effect
func (r *RouteStat) Start() error {
	if r.stats != nil {
		return r.stats.Start()
	}
	return errors.New("routestat: fail to start")
}

// Record stats module record the last node before
func (r *RouteStat) Record() error {
	if r.stats != nil {
		return r.stats.Record()
	}
	return errors.New("routestat: fail to record")
}

// Accept if stats module accept ip access than return true
func (r *RouteStat) Accept(ip string) bool {
	if r.throttler != nil && r.throttler.Accept(ip) {
		if r.stats != nil {
			return r.stats.Accept(ip)
		}
		return true
	}
	return false
}

// String debug or print out stats module information
func (r *RouteStat) String() string {
	if r.stats != nil {
		return r.stats.String()
	}
	return ""
}

// Close close stats module
func (r *RouteStat) Close() error {
	if r.stats != nil {
		if err := r.stats.Close(); err != nil {
			return err
		}
	}
	if r.throttler != nil {
		if err := r.throttler.Close(); err != nil {
			return err
		}
	}
	return nil
}
