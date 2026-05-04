package semaphore

import (
	"context"
	"sync"
)

type SemaphoreTwo struct {
	available int64
	mutex     sync.Mutex
	cond      *sync.Cond
	capacity  int64
}

func NewSemaphoreTwo(n int64) *SemaphoreTwo {
	mu := sync.Mutex{}

	return &SemaphoreTwo{
		available: n,
		cond:      sync.NewCond(&mu),
		capacity:  n,
	}
}

func (s *SemaphoreTwo) Acquire(ctx context.Context, n int64) error {
	s.mutex.Lock()
	for s.available < n {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			s.cond.Wait()
		}
	}
	s.available -= n

	s.mutex.Unlock()

	return nil
}

func (s *SemaphoreTwo) TryAcquire(n int64) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.available < n {
		return false
	}
	s.available -= n

	return true
}

func (s *SemaphoreTwo) Release(n int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.available+n > s.capacity {
		panic(ErrSemaphoreOverflow)
	}

	s.available += n
	s.cond.Broadcast()
}
