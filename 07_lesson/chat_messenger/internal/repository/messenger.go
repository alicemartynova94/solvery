package repository

import (
	"context"
	"github.com/google/uuid"
	"solvery/07_lesson/chat_messenger/internal/models"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, maxSessions int) (*models.Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID) error
	CheckSessionActive(id uuid.UUID) error
}

type ChatRepository interface {
	CreateChat(ctx context.Context, chat *models.Chat, maxChats int) (*models.Chat, error)
}
