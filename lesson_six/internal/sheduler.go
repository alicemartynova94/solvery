package internal

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID        uuid.UUID
	ExecuteAt time.Time     // для однократных
	Interval  time.Duration // для периодических (0 - однократная)
	Command   func(ctx context.Context) error
	Recurring bool
}

type Scheduler interface {
	ScheduleAt(Task, time.Time)
	ScheduleAfter(Task, time.Duration) // выполнить через N времени
	ScheduleEvery(Task, time.Duration) // периодическое выполнение
	Cancel(uuid.UUID)
	GetPendingTasks() []Task // получить список ожидающих задач
	Stop()                   // остановить планировщик (дождаться выполнения текущих)
}

type Schedule struct {
	m map[uuid.UUID]*Task
}

func (s *Schedule) NewScheduler(ctx context.Context) {

	go func(c context.Context) {
		for {

			for _, v := range s.m {
				if time.Now().After(v.ExecuteAt) {
					go v.Command(ctx)

					if v.Recurring {
						v.ExecuteAt = time.Now().Add(v.Interval)
					}

				}
			}

		}
	}(ctx)

}

func (s *Schedule) ScheduleAt(t Task, tm time.Time) {
	t.ExecuteAt = tm
	t.Interval = 0
	t.Recurring = false

	s.m[t.ID] = &t
}

func (s *Schedule) ScheduleAfter(t Task, d time.Duration) {
	t.ExecuteAt = time.Now().Add(d)
	t.Interval = 0
	t.Recurring = false

	s.m[t.ID] = &t
}

func (s *Schedule) ScheduleEvery(t Task, d time.Duration) {
	t.ExecuteAt = time.Now().Add(d)
	t.Interval = d
	t.Recurring = true

	s.m[t.ID] = &t
}

func (s *Schedule) Cancel(id uuid.UUID) {
	if _, ok := s.m[id]; ok {
		delete(s.m, id)
	}
}

func (s *Schedule) GetPendingTasks() []Task {
	tasks := make([]Task, 0, len(s.m))
	for _, v := range s.m {
		if time.Now().After(v.ExecuteAt) { //???
			tasks = append(tasks, *v)
		}
	}

	return tasks
}
