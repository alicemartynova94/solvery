package models

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ID          uuid.UUID
	ChatAuthor  Session
	ChatName    string
	ExpiresAt   *time.Time
	ReadOnly    bool
	MaxMessages int
	Members     []Session
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
