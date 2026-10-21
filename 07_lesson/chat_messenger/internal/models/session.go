package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID        uuid.UUID
	CreatedAt time.Time
	DeletedAt *time.Time
}
