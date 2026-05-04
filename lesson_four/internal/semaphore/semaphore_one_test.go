package semaphore

import (
	"context"
	"testing"
	"time"
)

func TestSemaphoreOne_Acquire_Success(t *testing.T) {
	s := NewSemaphoreOne(3)

	err := s.Acquire(context.Background(), 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSemaphoreOne_Acquire_Fail(t *testing.T) {
	s := NewSemaphoreOne(1)

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

func TestSemaphoreOne_Acquire_Context(t *testing.T) {
	s := NewSemaphoreOne(1)

	_ = s.Acquire(context.Background(), 1)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	err := s.Acquire(ctx, 1)

	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
}

func TestSemaphoreOne_Release_Success(t *testing.T) {
	s := NewSemaphoreOne(2)

	_ = s.Acquire(context.Background(), 2)

	s.Release(2)

	err := s.Acquire(context.Background(), 2)
	if err != nil {
		t.Fatal(err)
	}
}
