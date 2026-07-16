package internal_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	intrn "solvery/05_lesson/internal"
	"testing"
	"time"
)

func TestWorkpool_Run_Success(t *testing.T) {
	tests := []struct {
		name       string
		tasks      []intrn.Task
		workersNum int
		errNum     int
	}{
		{name: "no tasks for three workers",
			tasks:      []intrn.Task{},
			workersNum: 3,
			errNum:     1,
		},
		{name: "one task for three workers",
			tasks: []intrn.Task{
				func() error { return nil },
			},
			workersNum: 3,
			errNum:     1,
		},
		{name: "three tasks for three workers",
			tasks: []intrn.Task{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			workersNum: 3,
			errNum:     1,
		},
		{name: "five tasks for three workers",
			tasks: []intrn.Task{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			workersNum: 3,
			errNum:     1,
		},
		{name: "five tasks for three workers with errors",
			tasks: []intrn.Task{
				func() error { return nil },
				func() error { return errors.New("err") },
				func() error { return nil },
				func() error { return errors.New("err") },
				func() error { return nil },
			},
			workersNum: 3,
			errNum:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := intrn.Run(tt.tasks, tt.workersNum, tt.errNum)
			assert.NoError(t, err)
		})
	}
}

func TestWorkpool_Run_ReturnErr(t *testing.T) {
	tests := []struct {
		name       string
		tasks      []intrn.Task
		workersNum int
		errNum     int
	}{
		{name: " all tasks with errors",
			tasks: []intrn.Task{
				func() error { return errors.New("err") },
				func() error { return errors.New("err") },
				func() error { return errors.New("err") },
				func() error { return errors.New("err") },
				func() error { return errors.New("err") },
			},
			workersNum: 2,
			errNum:     3,
		},
		{name: " some tasks with errors exceeding limit",
			tasks: []intrn.Task{
				func() error { return errors.New("err") },
				func() error { return nil },
				func() error { return errors.New("err") },
				func() error { return nil },
				func() error { return errors.New("err") },
				func() error { return nil },
				func() error { return errors.New("err") },
			},
			workersNum: 4,
			errNum:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := intrn.Run(tt.tasks, tt.workersNum, tt.errNum)
			assert.ErrorIs(t, err, intrn.ErrErrorsLimitExceeded)
		})
	}
}

func TestWorkpool_Run_NoDeadLock(t *testing.T) {
	tasks := make([]intrn.Task, 10000)

	for i := range tasks {
		tasks[i] = func() error { return nil }
	}

	done := make(chan struct{})

	go func() {
		intrn.Run(tasks, 10, 3)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Error("timed out waiting for deadlock")
	}
}

func TestWorkpool_Run_ErrNumEqualZero(t *testing.T) {
	tests := []struct {
		name       string
		tasks      []intrn.Task
		workersNum int
		errNum     int
	}{
		{name: "zero err number",
			tasks: []intrn.Task{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			workersNum: 3,
			errNum:     0,
		},
		{name: "negative err number",
			tasks: []intrn.Task{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			workersNum: 3,
			errNum:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := intrn.Run(tt.tasks, tt.workersNum, tt.errNum)
			assert.NoError(t, err)
		})
	}
}
