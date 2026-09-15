package http

import (
	"net/http"
	"solvery/07_lesson/chat_messenger/internal/metrics"
	"solvery/07_lesson/chat_messenger/pkg/openapi"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateSession(w http.ResponseWriter, r *http.Request) {
	metrics.RequestsTotal.WithLabelValues(r.Method, "/sessions").Inc()

	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) GetChat(w http.ResponseWriter, r *http.Request, chatID openapi.ChatID) {
	metrics.RequestsTotal.WithLabelValues(r.Method, "/chat").Inc()
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
