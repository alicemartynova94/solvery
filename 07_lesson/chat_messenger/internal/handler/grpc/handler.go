package grpc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	dmnerr "solvery/07_lesson/chat_messenger/internal/errors"
	"solvery/07_lesson/chat_messenger/internal/service"
	"solvery/07_lesson/chat_messenger/pkg/messenger/v1"
)

type Handler struct {
	messenger.UnimplementedMessengerServiceServer

	sessionSrv service.SessionSrv
	logger     *zap.Logger
}

func NewHandler(sessionSrv service.SessionSrv, logger *zap.Logger) *Handler {
	return &Handler{
		sessionSrv: sessionSrv,
		logger:     logger,
	}
}

func (h *Handler) CreateSession(ctx context.Context, req *messenger.CreateSessionRequest) (*messenger.CreateSessionResponse, error) {
	session, err := h.sessionSrv.CreateSession(ctx)
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, "request canceled")
		case errors.Is(err, dmnerr.ErrSessionsLimitExceeded):
			return nil, status.Error(codes.ResourceExhausted, "sessions limit exceeded")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &messenger.CreateSessionResponse{
		Session: &messenger.Session{
			Id: session.ID.String(),
		},
	}, nil
}

func (h *Handler) DeleteSession(ctx context.Context, req *messenger.DeleteSessionRequest) (*messenger.DeleteSessionResponse, error) {
	if req == nil || req.Session == nil {
		return nil, status.Error(codes.InvalidArgument, "session is required")
	}

	id, err := uuid.Parse(req.Session.Id)
	if err != nil || id == uuid.Nil {
		return nil, status.Error(codes.InvalidArgument, "invalid session ID")
	}

	if err := h.sessionSrv.DeleteSession(ctx, id); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, "request canceled")
		case errors.Is(err, dmnerr.ErrSessionNotFound):
			return nil, status.Error(codes.NotFound, "session not found")
		case errors.Is(err, dmnerr.ErrSessionDeleted):
			return nil, status.Error(codes.FailedPrecondition, "session already deleted")
		case errors.Is(err, dmnerr.ErrSessionOwnsActiveChat):
			return nil, status.Error(codes.FailedPrecondition, "session owns active chat")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}
	return &messenger.DeleteSessionResponse{}, nil
}

func (*Handler) CreateChat(ctx context.Context, req *messenger.CreateChatRequest) (*messenger.CreateChatResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) GetChat(ctx context.Context, req *messenger.GetChatRequest) (*messenger.Chat, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) UpdateChat(ctx context.Context, req *messenger.UpdateChatRequest) (*messenger.Chat, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) DeleteChat(ctx context.Context, req *messenger.DeleteChatRequest) (*messenger.DeleteChatResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) AddChatMembers(ctx context.Context, req *messenger.AddChatMembersRequest) (*messenger.AddChatMembersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) DeleteChatMember(ctx context.Context, req *messenger.DeleteChatMemberRequest) (*messenger.DeleteChatMemberResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) CreateMessage(ctx context.Context, req *messenger.CreateMessageRequest) (*messenger.Message, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) GetMessages(ctx context.Context, req *messenger.GetMessagesRequest) (*messenger.GetMessagesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) DeleteMessage(ctx context.Context, req *messenger.DeleteMessageRequest) (*messenger.DeleteMessageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) EditMessage(ctx context.Context, req *messenger.EditMessageRequest) (*messenger.Message, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
