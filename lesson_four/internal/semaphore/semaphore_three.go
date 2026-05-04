package semaphore

import (
	"context"
	"sync/atomic"
)

type SemaphoreThree struct {
	available int64
	capacity  int64
}

func NewSemaphoreThree(n int64) *SemaphoreThree {
	return &SemaphoreThree{
		available: n,
		capacity:  n,
	}
}

func (s *SemaphoreThree) Acquire(ctx context.Context, n int64) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}

		val := atomic.LoadInt64(&s.available)

		if val < n {
			continue
		}

		newVal := val - n

		if atomic.CompareAndSwapInt64(&s.available, val, newVal) {
			return nil
		}
	}
}

func (s *SemaphoreThree) TryAcquire(n int64) bool {
	val := atomic.LoadInt64(&s.available)

	if val >= n {

		newVal := val - n

		if atomic.CompareAndSwapInt64(&s.available, val, newVal) {
			return true
		}
	}
	return false
}

func (s *SemaphoreThree) Release(n int64) {
	for {
		val := atomic.LoadInt64(&s.available)
		newVal := val + n

		if newVal > s.capacity {
			panic(ErrSemaphoreOverflow)
		}

		if atomic.CompareAndSwapInt64(&s.available, val, newVal) {
			return
		}
	}
}
