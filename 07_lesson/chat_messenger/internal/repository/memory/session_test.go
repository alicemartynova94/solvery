package memory_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	dmnerr "solvery/07_lesson/chat_messenger/internal/errors"
	"solvery/07_lesson/chat_messenger/internal/repository/memory"
	"solvery/07_lesson/chat_messenger/internal/repository/synchronization"
	"testing"
)

func newTestRepository() *memory.SessionRepository {
	return memory.NewSessionRepository(synchronization.NewMutex())
}

func TestCreateSession(t *testing.T) {
	tests := []struct {
		name        string
		maxSessions int
		existing    int
		wantErr     error
	}{
		{
			name:        "create session without limit",
			maxSessions: 0,
			existing:    10,
		},
		{
			name:        "create session within limit",
			maxSessions: 3,
			existing:    2,
		},
		{
			name:        "session limit exceeded",
			maxSessions: 1,
			existing:    1,
			wantErr:     dmnerr.ErrSessionsLimitExceeded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newTestRepository()
			ctx := context.Background()

			for range tt.existing {
				if _, err := repo.CreateSession(ctx, 0); err != nil {
					require.NoError(t, err)
				}
			}

			session, err := repo.CreateSession(ctx, tt.maxSessions)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("CreateSession() error = %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, session)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, session)

			assert.NotEqual(t, uuid.Nil, session.ID)
			assert.False(t, session.CreatedAt.IsZero())
			assert.Nil(t, session.DeletedAt)
		})
	}
}

func TestDeleteSession(t *testing.T) {
	tests := []struct {
		name       string
		create     bool
		deleteOnce bool
		wantErr    error
		wantDelete bool
	}{
		{
			name:       "delete active session",
			create:     true,
			wantDelete: true,
		},
		{
			name:    "delete nonexistent session",
			wantErr: dmnerr.ErrSessionNotFound,
		},
		{
			name:       "delete already deleted session",
			create:     true,
			deleteOnce: true,
			wantErr:    dmnerr.ErrSessionDeleted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newTestRepository()
			ctx := context.Background()

			var sessionID uuid.UUID

			if tt.create {
				session, err := repo.CreateSession(ctx, 0)
				require.NoError(t, err)
				require.NotNil(t, session)

				sessionID = session.ID
			} else {
				sessionID = uuid.New()
			}

			if tt.deleteOnce {
				err := repo.DeleteSession(ctx, sessionID)
				require.NoError(t, err)
			}

			err := repo.DeleteSession(ctx, sessionID)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)

			if tt.wantDelete {
				err := repo.CheckSessionActive(sessionID)
				assert.ErrorIs(t, err, dmnerr.ErrSessionDeleted)
			}
		})
	}
}
