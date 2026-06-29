package internal

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
	scheduler := NewSchedule()

}
