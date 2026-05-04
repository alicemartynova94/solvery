package internal

import "sync"

func or(channels ...<-chan interface{}) <-chan interface{} {
	var o sync.Once
	done := make(chan interface{})
	for _, ch := range channels {
		go func(c <-chan interface{}) {
			for range c {
			}

			o.Do(func() {
				close(done)
			})

		}(ch)
	}
	return done
}
