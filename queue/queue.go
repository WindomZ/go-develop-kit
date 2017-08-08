package queue

import "sync"

// Queue defines a basic queue model.
type Queue interface {
	Capacity() int
	Size() int
	Push(i interface{}) int
	PushSlice(s []interface{}) int
	IsIdle() bool
	Pull() interface{}
	PullSync() interface{}
	Exchange(size uint) bool
	Free()
}

// Pool the Queue instance.
type Pool struct {
	Cap  int
	lock sync.Mutex
	Chan chan interface{}
}

// Capacity returns integer, capacity of queue.
func (p Pool) Capacity() int {
	return p.Cap
}

// Size returns integer, length of queue.
func (p Pool) Size() int {
	return len(p.Chan)
}

// Push returns integer, length of queue.
// Push a item to the queue.
func (p *Pool) Push(i interface{}) int {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.Chan) >= p.Cap {
		return -1
	}

	p.Chan <- i
	return len(p.Chan)
}

// PushSlice returns integer, length of queue.
// Push slice of items to the queue.
func (p *Pool) PushSlice(s []interface{}) int {
	if len(s) == 0 {
		return len(p.Chan)
	}
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.Chan)+len(s) > p.Cap {
		return -1
	}

	for _, i := range s {
		p.Chan <- i
	}

	return len(p.Chan)
}

// IsIdle returns true if queue is empty
func (p Pool) IsIdle() bool {
	return len(p.Chan) == 0
}

// Pull returns the first item of queue
func (p *Pool) Pull() interface{} {
	return <-p.Chan
}

// PullSync returns the first item of queue, nil if queue is empty.
func (p *Pool) PullSync() interface{} {
	if len(p.Chan) == 0 {
		return nil
	}

	return <-p.Chan
}

// Exchange returns boolean, if exchanges queue capacity with size successfully.
// returns fail if length of queue large than size.
func (p *Pool) Exchange(size uint) bool {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.Chan) >= int(size) {
		return false
	}

	if len(p.Chan) >= 1 {
		pool := make([]interface{}, 0, len(p.Chan))
		for len(p.Chan) != 0 {
			pool = append(pool, <-p.Chan)
		}
		p.Cap = int(size)
		p.Chan = make(chan interface{}, p.Cap)
		for _, i := range pool {
			p.Chan <- i
		}
	} else {
		p.Cap = int(size)
		p.Chan = make(chan interface{}, p.Cap)
	}

	return true
}

// Free release all queue items.
func (p *Pool) Free() {
	p.lock.Lock()
	for len(p.Chan) != 0 {
		<-p.Chan
	}
	p.lock.Unlock()
}

// New returns a Queue interface.
func New(size uint) Queue {
	return &Pool{
		Cap:  int(size),
		Chan: make(chan interface{}, int(size)),
	}
}
