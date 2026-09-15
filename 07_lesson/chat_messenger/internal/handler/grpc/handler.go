package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"solvery/07_lesson/chat_messenger/pkg/messenger/v1"
)

type Handler struct {
	messenger.UnimplementedMessengerServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) CreateSession(ctx context.Context, req *messenger.CreateSessionRequest) (*messenger.Session, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) DeleteSession(ctx context.Context, req *emptypb.Empty) (*messenger.DeleteSessionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Handler) CreateChat(ctx context.Context, req *messenger.CreateChatRequest) (*messenger.Chat, error) {
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
