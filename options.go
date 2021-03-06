package sync

import (
	"github.com/micro/go-sync/data"
	"github.com/micro/go-sync/leader"
	"github.com/micro/go-sync/lock"
	"github.com/micro/go-sync/time"
)

// WithLeader sets the leader election implementation opton
func WithLeader(l leader.Leader) Option {
	return func(o *Options) {
		o.Leader = l
	}
}

// WithLock sets the locking implementation option
func WithLock(l lock.Lock) Option {
	return func(o *Options) {
		o.Lock = l
	}
}

// WithData sets the data implementation option
func WithData(s data.Data) Option {
	return func(o *Options) {
		o.Data = s
	}
}

// WithTime sets the time implementation option
func WithTime(t time.Time) Option {
	return func(o *Options) {
		o.Time = t
	}
}
