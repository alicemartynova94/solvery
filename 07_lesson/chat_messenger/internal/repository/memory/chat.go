package memory

import (
	"context"
	"github.com/google/uuid"
	dmnerr "solvery/07_lesson/chat_messenger/internal/errors"
	"solvery/07_lesson/chat_messenger/internal/models"
)

type ChatInMemoryRepository struct {
	chats map[uuid.UUID]models.Chat
}

func NewChatRepository() *ChatInMemoryRepository {
	return &ChatInMemoryRepository{
		chats: make(map[uuid.UUID]models.Chat),
	}
}

func (r *ChatInMemoryRepository) CreateChat(ctx context.Context, chat *models.Chat, maxChats int) (*models.Chat, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	activeChats := 0
	for _, v := range r.chats {
		if v.DeletedAt == nil {
			activeChats++
		}
	}

	if maxChats != 0 && activeChats >= maxChats {
		return nil, dmnerr.ErrChatLimitExceeded
	}

	r.chats[chat.ID] = *chat

	return chat, nil
}
