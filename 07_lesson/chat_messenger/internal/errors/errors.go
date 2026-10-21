package errors

import "errors"

var (
	ErrInvalidUUIDFormat = errors.New("id is either blank or has invalid format")

	ErrSessionNotFound       = errors.New("session not found")
	ErrSessionDeleted        = errors.New("session was already deleted")
	ErrSessionsLimitExceeded = errors.New("session limit exceeded")
	ErrSessionOwnsActiveChat = errors.New("session owns active chat and cannot be deleted")

	ErrChatLimitExceeded       = errors.New("chat limit exceeded")
	ErrChatInvalidExpiredAt    = errors.New("invalid chat expired time")
	ErrChatNameEmpty           = errors.New("chat_name must not be empty")
	ErrChatNameExceedLimit     = errors.New("chat_name must be at most 100 characters")
	ErrChatNegativeMaxMessages = errors.New("max_messages must be non-negative")
)
