package internal

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Task struct {
	ID        uuid.UUID
	ExecuteAt time.Time
	Interval  time.Duration
	Command   func(ctx context.Context) error
	Recurring bool
}

type Scheduler interface {
	ScheduleAt(Task, time.Time)
	ScheduleAfter(Task, time.Duration)
	ScheduleEvery(Task, time.Duration)
	Cancel(uuid.UUID)
	GetPendingTasks() []Task
	Stop()
}

type Schedule struct {
	mu     *sync.RWMutex
	wg     *sync.WaitGroup
	m      map[uuid.UUID]*Task
	cancel context.CancelFunc
}

func NewSchedule() *Schedule {
	return &Schedule{
		mu: &sync.RWMutex{},
		wg: &sync.WaitGroup{},
		m:  make(map[uuid.UUID]*Task),
	}
}

func (s *Schedule) ScheduleAt(t Task, tm time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.ExecuteAt = tm
	t.Interval = 0
	t.Recurring = false

	s.m[t.ID] = &t
}

func (s *Schedule) ScheduleAfter(t Task, d time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.ExecuteAt = time.Now().Add(d)
	t.Interval = 0
	t.Recurring = false

	s.m[t.ID] = &t
}

func (s *Schedule) ScheduleEvery(t Task, d time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.ExecuteAt = time.Now().Add(d)
	t.Interval = d
	t.Recurring = true

	s.m[t.ID] = &t
}

func (s *Schedule) Cancel(id uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.m[id]; ok {
		delete(s.m, id)
	}
}

func (s *Schedule) GetPendingTasks() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]Task, 0, len(s.m))
	for _, v := range s.m {
		if time.Now().Before(v.ExecuteAt) {
			tasks = append(tasks, *v)
		}
	}

	return tasks
}

func (s *Schedule) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
	s.wg.Wait()
}

func (s *Schedule) Run(ctx context.Context) {
	localContext, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	go func() {

		timer := time.NewTicker(100 * time.Millisecond)
		defer timer.Stop()
		for {

			select {
			case <-localContext.Done():
				return
			case <-timer.C:

				var tasks []*Task

				s.mu.Lock()

				for _, v := range s.m {
					if time.Now().After(v.ExecuteAt) {
						tasks = append(tasks, v)

						if v.Recurring {
							v.ExecuteAt = v.ExecuteAt.Add(v.Interval)
						} else {
							delete(s.m, v.ID)
						}

					}
				}
				s.mu.Unlock()

				for _, v := range tasks {
					s.wg.Add(1)

					go func(t *Task) {
						defer s.wg.Done()
						_ = t.Command(localContext)
					}(v)
				}
			}
		}
	}()

}
