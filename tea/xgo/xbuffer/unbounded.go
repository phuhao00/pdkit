package xbuffer

import "sync"

type Unbounded struct {
	c       chan interface{}
	mu      sync.Mutex
	backlog []interface{}
}

// NewUnbounded returns a new instance of Unbounded.
func NewUnbounded() *Unbounded {
	return &Unbounded{c: make(chan interface{}, 1)}
}

// Put adds t to the unbounded buffer.
func (b *Unbounded) Put(t interface{}) {
	b.mu.Lock()
	if len(b.backlog) == 0 {
		select {
		case b.c <- t:
			b.mu.Unlock()
			return
		default:
		}
	}
	b.backlog = append(b.backlog, t)
	b.mu.Unlock()
}

// Load sends the earliest buffered data, if any, onto the read channel
// returned by Get(). Users are expected to call this every time they read a
// value from the read channel.
func (b *Unbounded) Load() {
	b.mu.Lock()
	if len(b.backlog) > 0 {
		select {
		case b.c <- b.backlog[0]:
			b.backlog[0] = nil
			b.backlog = b.backlog[1:]
		default:
		}
	}
	b.mu.Unlock()
}

// Get returns a read channel on which values added to the buffer, via Put(),
// are sent on.
//
// Upon reading a value from this channel, users are expected to call Load() to
// send the next buffered value onto the channel if there is any.
func (b *Unbounded) Get() <-chan interface{} {
	return b.c
}
