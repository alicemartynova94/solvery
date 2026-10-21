-- +goose Up
CREATE TABLE messages
(
    id         UUID PRIMARY KEY         NOT NULL DEFAULT uuid_generate_v4(),
    chat_id    UUID                     NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    author_id  UUID                     NOT NULL REFERENCES sessions (id),
    text       TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- +goose Down
DROP TABLE IF EXISTS messages;
