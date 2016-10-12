package routestat

import (
	"github.com/tsenart/tb"
	"time"
)

type Throttler struct {
	defaultRate int64
	throttler   *tb.Throttler
}

func NewThrottler(freq time.Duration) *Throttler {
	return &Throttler{
		defaultRate: int64(time.Second / freq),
		throttler:   tb.NewThrottler(freq),
	}
}

func (t *Throttler) Start() error {
	return nil
}

func (t *Throttler) Record() error {
	return nil
}

func (t *Throttler) Accept(ip string) bool {
	return !t.throttler.Halt(ip, 1, t.defaultRate)
}

func (t *Throttler) String() string {
	return ""
}

func (t *Throttler) Close() error {
	return t.throttler.Close()
}
