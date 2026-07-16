package internal_test

import (
	intnl "solvery/04_lesson/internal"
	"testing"
)

func TestOr(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	done := intnl.Or(ch1, ch2, ch3)

	close(ch2)

	_, ok := <-done

	if ok != false {
		t.Error("array of channels was not closed")
	}
}
