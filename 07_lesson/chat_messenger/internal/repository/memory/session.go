package memory

import (
	"context"
	"github.com/google/uuid"
	dmnerr "solvery/07_lesson/chat_messenger/internal/errors"
	"solvery/07_lesson/chat_messenger/internal/models"
	"solvery/07_lesson/chat_messenger/internal/repository/synchronization"
	"time"
)

type SessionRepository struct {
	sessions map[uuid.UUID]models.Session
	mu       *synchronization.Mutex
}

func NewSessionRepository(mu *synchronization.Mutex) *SessionRepository {
	return &SessionRepository{
		sessions: make(map[uuid.UUID]models.Session),
		mu:       mu,
	}
}

func (r *SessionRepository) CreateSession(ctx context.Context, maxSessions int) (*models.Session, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	activeSessions := 0
	for _, v := range r.sessions {
		if v.DeletedAt == nil {
			activeSessions++
		}
	}

	if maxSessions != 0 && activeSessions >= maxSessions {
		return nil, dmnerr.ErrSessionsLimitExceeded
	}

	newSession := models.Session{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}

	r.sessions[newSession.ID] = newSession

	return &newSession, nil
}

func (r *SessionRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return err
	}

	if err := r.CheckSessionActive(id); err != nil {
		return err
	}

	deletedAt := time.Now()
	session := r.sessions[id]
	session.DeletedAt = &deletedAt
	r.sessions[id] = session

	return nil
}

func (r *SessionRepository) CheckSessionActive(id uuid.UUID) error {
	session, ok := r.sessions[id]
	if !ok {
		return dmnerr.ErrSessionNotFound
	}

	if session.DeletedAt != nil {
		return dmnerr.ErrSessionDeleted
	}

	return nil
}
