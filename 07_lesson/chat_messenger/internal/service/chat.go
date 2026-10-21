package service

import (
	"context"
	"github.com/google/uuid"
	dmnerr "solvery/07_lesson/chat_messenger/internal/errors"
	"solvery/07_lesson/chat_messenger/internal/handler/dto"
	"solvery/07_lesson/chat_messenger/internal/mapper"
	"solvery/07_lesson/chat_messenger/internal/models"
	"solvery/07_lesson/chat_messenger/internal/repository"
	"solvery/07_lesson/chat_messenger/internal/repository/synchronization"
	"time"
)

type ChatSrv interface {
	CreateChat(ctx context.Context, reqDTO *dto.CreateChatRequest) (*dto.CreateChatResponse, error)
	GetChat(ctx context.Context, id uuid.UUID) (*models.Chat, error)
}

type ChatService struct {
	mu          *synchronization.Mutex
	chatRepo    repository.ChatRepository
	sessionRepo repository.SessionRepository
	maxChats    int
}

func NewChatService(mu *synchronization.Mutex, chatRepo repository.ChatRepository, sessionRepo repository.SessionRepository, maxChats int) *ChatService {
	return &ChatService{
		mu:          mu,
		chatRepo:    chatRepo,
		sessionRepo: sessionRepo,
		maxChats:    maxChats,
	}
}

func (s *ChatService) CreateChat(ctx context.Context, reqDTO *dto.CreateChatRequest) (*dto.CreateChatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.sessionRepo.CheckSessionActive(reqDTO.ChatAuthor)
	if err != nil {
		return nil, err
	}

	if reqDTO.ExpiresAt != nil {
		if reqDTO.ExpiresAt.Before(time.Now()) {
			return nil, dmnerr.ErrChatInvalidExpiredAt
		}
	}

	newChat, err := s.chatRepo.CreateChat(ctx, mapper.CreateChatDTOtoChat(reqDTO), s.maxChats)
	if err != nil {
		return nil, err
	}

	return mapper.ChatToCreateChatRespDTO(newChat), nil
}

func (s *ChatService) GetChat(ctx context.Context, id uuid.UUID) (*models.Chat, error) {
	return s.chatRepo.GetChat(ctx, id)
}
