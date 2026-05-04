package semaphore

import (
	"context"
	"errors"
)

var ErrSemaphoreOverflow = errors.New("semaphore overflow: cannot release more than the capacity")

type Semaphore interface {
	Acquire(context.Context, int64) error
	TryAcquire(int64) bool
	Release(int64)
}
