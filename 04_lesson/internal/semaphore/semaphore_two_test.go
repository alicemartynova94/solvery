package semaphore_test

import (
	"context"
	"solvery/04_lesson/internal/semaphore"
	"testing"
	"time"
)

func TestSemaphoreTwo_Acquire_Success(t *testing.T) {
	s := semaphore.NewSemaphoreTwo(3)

	err := s.Acquire(context.Background(), 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSemaphoreTwo_Acquire_Fail(t *testing.T) {
	s := semaphore.NewSemaphoreTwo(1)

	ctx := context.Background()

	_ = s.Acquire(ctx, 1)

	done := make(chan struct{})

	go func() {
		err := s.Acquire(ctx, 1)
		if err == nil {
			t.Error("expected error, got nil")
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
	}
}

//func TestSemaphoreTwo_Acquire_Context(t *testing.T) {
//	s := NewSemaphoreTwo(0)
//
//	ctx, cancel := context.WithCancel(context.Background())
//
//	done := make(chan error)
//
//	go func() {
//		done <- s.Acquire(ctx, 1)
//	}()
//
//	time.Sleep(50 * time.Millisecond)
//	cancel()
//
//	err := <-done
//
//	if errors.Is(err, context.Canceled) {
//		t.Fatalf("expected context.Canceled, got %v", err)
//	}
//}

func TestSemaphoreTwo_Release_Success(t *testing.T) {
	s := semaphore.NewSemaphoreTwo(2)

	_ = s.Acquire(context.Background(), 2)

	s.Release(2)

	err := s.Acquire(context.Background(), 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSemaphoreTwo_TryAcquire(t *testing.T) {
	s := semaphore.NewSemaphoreTwo(2)

	tests := []struct {
		name     string
		n        int64
		expected bool
	}{
		{"first acquire", 1, true},
		{"second acquire", 1, true},
		{"acquire fail", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.TryAcquire(tt.n)
			if got != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
