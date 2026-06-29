package internal

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestScheduler_ScheduleAt(t *testing.T) {
	task := Task{
		ID: uuid.New(),
	}

	tm := time.Date(2026, 6, 30, 12, 0, 0, 0, time.UTC)
	schedule := NewSchedule()
	schedule.ScheduleAt(task, tm)

	el := schedule.m[task.ID]

	assert.Equal(t, tm, el.ExecuteAt)
	assert.Equal(t, time.Duration(0), el.Interval)
	assert.False(t, el.Recurring)
}

func TestScheduler_ScheduleAfter(t *testing.T) {
	task := Task{
		ID: uuid.New(),
	}

	tm := time.Now()
	duration := 30 * time.Minute
	schedule := NewSchedule()
	schedule.ScheduleAfter(task, duration)

	el := schedule.m[task.ID]

	assert.WithinDuration(t, tm.Add(duration), el.ExecuteAt, 50*time.Millisecond)
	assert.Equal(t, time.Duration(0), el.Interval)
	assert.False(t, el.Recurring)
}

func TestScheduler_ScheduleEvery(t *testing.T) {
	task := Task{
		ID: uuid.New(),
	}

	tm := time.Now()
	duration := 30 * time.Minute
	schedule := NewSchedule()
	schedule.ScheduleEvery(task, duration)

	el := schedule.m[task.ID]

	assert.WithinDuration(t, tm.Add(duration), el.ExecuteAt, 50*time.Millisecond)
	assert.Equal(t, duration, el.Interval)
	assert.True(t, el.Recurring)
}

func TestScheduler_CancelAndGetPending(t *testing.T) {
	tm := time.Now().Add(30 * time.Minute)
	task1 := Task{
		ID: uuid.New(),
	}
	task2 := Task{
		ID: uuid.New(),
	}

	schedule := NewSchedule()
	schedule.ScheduleAt(task1, tm)
	schedule.ScheduleAt(task2, tm)

	schedule.Cancel(task2.ID)

	tasks := schedule.GetPendingTasks()

	for _, task := range tasks {
		assert.NotEqual(t, task.ID, task2.ID)
	}
}

func TestScheduler_RunStop(t *testing.T) {
	ctx := context.Background()
	scheduler := NewSchedule()
	done := make(chan struct{})
	scheduler.Run(ctx)

	task := Task{
		ID: uuid.New(),
		Command: func(ctx context.Context) error {
			close(done)
			return nil
		},
	}

	scheduler.ScheduleAfter(task, 5*time.Millisecond)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("task was not performed")
	}

	scheduler.Stop()
}

func TestScheduler_Concurrent(t *testing.T) {
	scheduler := NewSchedule()

	ctx := context.Background()
	scheduler.Run(ctx)

	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := uuid.New()

			scheduler.ScheduleAfter(Task{
				ID: id,
				Command: func(ctx context.Context) error {
					return nil
				},
			}, 10*time.Millisecond)

			scheduler.GetPendingTasks()
			scheduler.Cancel(id)
		}()
	}

	wg.Wait()
	scheduler.Stop()
}
