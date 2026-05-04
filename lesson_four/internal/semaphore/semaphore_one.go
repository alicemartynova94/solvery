package semaphore

import (
	"context"
)

type SemaphoreOne struct {
	ch chan struct{}
}

func NewSemaphoreOne(cap uint) *SemaphoreOne {
	newChannel := make(chan struct{}, cap)
	for i := 0; i < int(cap); i++ {
		newChannel <- struct{}{}
	}
	return &SemaphoreOne{ch: newChannel}
}

func (s *SemaphoreOne) Acquire(ctx context.Context, n int64) error {
	available := 0

	for i := 0; i < int(n); i++ {
		select {
		case <-ctx.Done():
			for j := 0; j < available; j++ {
				s.ch <- struct{}{}
			}
			return ctx.Err()

		case <-s.ch:
			available++
		}
	}

	return nil
}

func (s *SemaphoreOne) TryAcquire(n int64) bool {
	return true
}

func (s *SemaphoreOne) Release(n int64) {
	for i := 0; i < int(n); i++ {
		select {
		case s.ch <- struct{}{}:
		default:
			panic(ErrSemaphoreOverflow)
		}
	}
}
