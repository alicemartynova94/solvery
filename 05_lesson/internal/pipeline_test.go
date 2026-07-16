package internal_test

import (
	"reflect"
	intrn "solvery/05_lesson/internal"
	"testing"
	"time"
)

func TestPipeline_ExecutePipeline_Successful(t *testing.T) {
	in := make(chan interface{})
	done := make(chan interface{})

	stage1 := func(in intrn.In) intrn.Out {
		out := make(chan interface{})

		go func() {
			defer close(out)

			for v := range in {
				out <- v.(int) * 2
			}
		}()
		return out
	}

	stage2 := func(in intrn.In) intrn.Out {
		out := make(chan interface{})

		go func() {
			defer close(out)

			for v := range in {
				out <- v.(int) + 1
			}
		}()
		return out
	}

	out := intrn.ExecutePipeline(in, done, stage1, stage2)

	go func() {
		defer close(in)

		for i := 1; i < 4; i++ {
			in <- i
		}
	}()

	result := make([]int, 0, 3)

	for val := range out {
		result = append(result, val.(int))
	}

	expected := []int{3, 5, 7}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

func TestPipeline_ExecutePipelineWithDone(t *testing.T) {
	in := make(chan interface{})
	done := make(chan interface{})
	signal := make(chan struct{})

	stage1 := func(in intrn.In) intrn.Out {
		out := make(chan interface{})

		go func() {
			defer close(out)

			for v := range in {
				close(signal)
				out <- v.(int) * 2
			}
		}()
		return out
	}

	out := intrn.ExecutePipeline(in, done, stage1)

	go func() {
		defer close(in)

		for i := 1; i < 4; i++ {
			in <- i
		}
	}()

	<-signal
	close(done)

	select {
	case _, ok := <-out:
		if ok {
			t.Error("expected closed channel")
		}
	case <-time.After(time.Second):
		t.Error("pipeline did not complete on done")
	}
}

func TestPipeline_OutputClosed(t *testing.T) {
	in := make(chan interface{})
	done := make(chan interface{})

	out := intrn.ExecutePipeline(in, done)

	close(in)

	select {
	case _, ok := <-out:
		if ok {
			t.Error("expected closed output")
		}
	case <-time.After(time.Second):
		t.Error("timeout waiting closed output")
	}
}
