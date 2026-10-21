package service

import (
	"context"
	"github.com/google/uuid"
	"solvery/07_lesson/chat_messenger/internal/models"
	"solvery/07_lesson/chat_messenger/internal/repository"
)

type SessionSrv interface {
	CreateSession(ctx context.Context) (*models.Session, error)
	DeleteSession(ctx context.Context, ID uuid.UUID) error
}

type SessionService struct {
	repo        repository.SessionRepository
	maxSessions int
}

func NewSessionService(repo repository.SessionRepository, maxSessions int) *SessionService {
	return &SessionService{
		repo:        repo,
		maxSessions: maxSessions,
	}
}

func (s *SessionService) CreateSession(ctx context.Context) (*models.Session, error) {
	return s.repo.CreateSession(ctx, s.maxSessions)
}

// TODO: session owns active chat check
func (s *SessionService) DeleteSession(ctx context.Context, ID uuid.UUID) error {
	return s.repo.DeleteSession(ctx, ID)
}
