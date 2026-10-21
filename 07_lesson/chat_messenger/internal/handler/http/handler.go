package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	dmnerr "solvery/07_lesson/chat_messenger/internal/errors"
	"solvery/07_lesson/chat_messenger/internal/handler/dto"
	"solvery/07_lesson/chat_messenger/internal/service"
	"solvery/07_lesson/chat_messenger/pkg/openapi"
)

type Handler struct {
	sessionSrv service.SessionSrv
	chatSrv    service.ChatSrv
	logger     *zap.Logger
}

func NewHandler(sessionSrv service.SessionSrv, chatSrv service.ChatSrv, logger *zap.Logger) *Handler {
	return &Handler{
		sessionSrv: sessionSrv,
		chatSrv:    chatSrv,
		logger:     logger,
	}
}

func (h *Handler) CreateSession(w http.ResponseWriter, r *http.Request) {
	session, err := h.sessionSrv.CreateSession(r.Context())
	if err != nil {
		h.writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(session.ID); err != nil {
		h.logger.Error("failed to encode create session response", zap.Error(err))
		return
	}
}

func (h *Handler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionIDString := r.Header.Get("X-Session-ID")
	sessionID, err := parseID(sessionIDString)
	if err != nil {
		h.writeError(w, err)
		return
	}

	if err := h.sessionSrv.DeleteSession(r.Context(), sessionID); err != nil {
		h.writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	sessionIDString := r.Header.Get("X-Session-ID")
	sessionID, err := parseID(sessionIDString)
	if err != nil {
		h.writeError(w, err)
		return
	}

	var reqDTO dto.CreateChatRequest
	if err := json.NewDecoder(r.Body).Decode(&reqDTO); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := reqDTO.Validate(); err != nil {
		h.writeError(w, err)
		return
	}
	for _, memberId := range reqDTO.Members {
		if memberId == uuid.Nil {
			http.Error(w, "invalid member session id", http.StatusBadRequest)
			return
		}
	}
	reqDTO.ChatAuthor = sessionID

	createdChat, err := h.chatSrv.CreateChat(r.Context(), &reqDTO)
	if err != nil {
		h.writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdChat); err != nil {
		h.logger.Error("failed to encode create chat response", zap.Error(err))
		return
	}
}

func (h *Handler) GetChat(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) UpdateChat(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) DeleteChat(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) AddChatMembers(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) DeleteChatMember(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID, params openapi.GetMessagesParams) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID, messageID openapi.MessageID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) EditMessage(w http.ResponseWriter, r *http.Request, chatId openapi.ChatID, messageId openapi.MessageID) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, context.Canceled):
		return
	case errors.Is(err, dmnerr.ErrSessionsLimitExceeded):
		http.Error(w, "sessions limit exceeded", http.StatusConflict)
	case errors.Is(err, dmnerr.ErrInvalidUUIDFormat):
		http.Error(w, "missing X-Session-ID header or invalid format", http.StatusBadRequest)
	case errors.Is(err, dmnerr.ErrSessionNotFound):
		http.Error(w, "session not found", http.StatusNotFound)
	case errors.Is(err, dmnerr.ErrSessionDeleted):
		http.Error(w, "session was already deleted", http.StatusConflict)
	case errors.Is(err, dmnerr.ErrSessionOwnsActiveChat):
		http.Error(w, "session owns active chat", http.StatusConflict)
	case errors.Is(err, dmnerr.ErrChatLimitExceeded):
		http.Error(w, "chat limit exceeded", http.StatusConflict)
	case errors.Is(err, dmnerr.ErrChatInvalidExpiredAt):
		http.Error(w, "chat invalid expired at time", http.StatusBadRequest)
	default:
		h.logger.Error("failed to perform operation", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func parseID(sessionID string) (uuid.UUID, error) {
	if sessionID == "" {
		return uuid.Nil, dmnerr.ErrInvalidUUIDFormat
	}

	id, err := uuid.Parse(sessionID)
	if err != nil || id == uuid.Nil {
		return uuid.Nil, dmnerr.ErrInvalidUUIDFormat
	}

	return id, nil
}
