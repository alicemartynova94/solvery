package mapper

import (
	"github.com/google/uuid"
	"solvery/07_lesson/chat_messenger/internal/handler/dto"
	"solvery/07_lesson/chat_messenger/internal/models"
	"time"
)

func CreateChatDTOtoChat(dto *dto.CreateChatRequest) *models.Chat {
	members := make([]models.Session, 0, len(dto.Members))
	for _, v := range dto.Members {
		members = append(members, models.Session{ID: v})
	}
	members = append(members, models.Session{ID: dto.ChatAuthor})

	return &models.Chat{
		ID:          uuid.New(),
		ChatAuthor:  models.Session{ID: dto.ChatAuthor},
		ChatName:    dto.ChatName,
		ExpiresAt:   dto.ExpiresAt,
		ReadOnly:    dto.ReadOnly,
		MaxMessages: dto.MaxMessages,
		Members:     members,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ChatToCreateChatRespDTO(chat *models.Chat) *dto.CreateChatResponse {
	members := make([]uuid.UUID, 0, len(chat.Members))
	for _, v := range chat.Members {
		members = append(members, v.ID)
	}

	return &dto.CreateChatResponse{
		ID:          chat.ID,
		ChatAuthor:  chat.ChatAuthor.ID,
		ChatName:    chat.ChatName,
		ExpiresAt:   chat.ExpiresAt,
		ReadOnly:    chat.ReadOnly,
		MaxMessages: chat.MaxMessages,
		CreatedAt:   chat.CreatedAt,
		Members:     members,
	}
}
