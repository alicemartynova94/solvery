package dto

import (
	"github.com/google/uuid"
	dmnerr "solvery/07_lesson/chat_messenger/internal/errors"
	"strings"
	"time"
	"unicode/utf8"
)

type CreateChatRequest struct {
	ChatAuthor  uuid.UUID   `json:"-"`
	ChatName    string      `json:"chat_name"`
	ExpiresAt   *time.Time  `json:"expires_at"`
	ReadOnly    bool        `json:"read_only"`
	MaxMessages int         `json:"max_messages"`
	Members     []uuid.UUID `json:"members"`
}

type CreateChatResponse struct {
	ID          uuid.UUID   `json:"id"`
	ChatAuthor  uuid.UUID   `json:"chat_author"`
	ChatName    string      `json:"chat_name"`
	ExpiresAt   *time.Time  `json:"expires_at"`
	ReadOnly    bool        `json:"read_only"`
	MaxMessages int         `json:"max_messages"`
	CreatedAt   time.Time   `json:"created_at"`
	Members     []uuid.UUID `json:"members"`
}

func (r *CreateChatRequest) Validate() error {
	name := strings.TrimSpace(r.ChatName)

	if name == "" {
		return dmnerr.ErrChatNameEmpty
	}

	if utf8.RuneCountInString(name) > 100 {
		return dmnerr.ErrChatNameExceedLimit
	}

	if r.MaxMessages < 0 {
		return dmnerr.ErrChatNegativeMaxMessages
	}

	return nil
}
