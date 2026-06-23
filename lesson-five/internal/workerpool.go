package internal

import (
	"errors"
	"math"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		m = math.MaxInt64
	}

	taskCh := make(chan Task, n)
	doneCh := make(chan struct{})

	var wg sync.WaitGroup
	var errNum int64
	var once sync.Once

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-doneCh:
					return
				case task, ok := <-taskCh:
					if !ok {
						return
					}

					if err := task(); err != nil {
						if atomic.AddInt64(&errNum, 1) >= int64(m) {
							once.Do(func() {
								close(doneCh)
							})
							return
						}
					}
				}
			}
		}()
	}

	go func() {
		defer close(taskCh)

		for _, task := range tasks {
			select {
			case <-doneCh:
				return
			case taskCh <- task:
			}
		}
	}()

	wg.Wait()

	if atomic.LoadInt64(&errNum) >= int64(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
