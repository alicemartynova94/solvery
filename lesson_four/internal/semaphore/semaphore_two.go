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
	s := &SemaphoreTwo{
		available: n,
		capacity:  n,
	}

	s.cond = sync.NewCond(&s.mutex)
	return s
}

func (s *SemaphoreTwo) Acquire(ctx context.Context, n int64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for s.available < n {
		done := make(chan struct{})

		go func() {
			s.cond.Wait()
			close(done)
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-done:
		}
	}

	s.available -= n
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
