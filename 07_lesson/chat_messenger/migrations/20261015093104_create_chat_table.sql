-- +goose Up
CREATE TABLE chats
(
    id           UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    chat_author  UUID         NOT NULL REFERENCES sessions (id),
    chat_name    VARCHAR(100) NOT NULL,
    expires_at   TIMESTAMP WITH TIME ZONE,
    read_only    BOOLEAN      NOT NULL DEFAULT FALSE,
    max_messages INTEGER      NOT NULL DEFAULT 0,
    created_at   TIMESTAMP WITH TIME ZONE  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE  NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMP WITH TIME ZONE
);

-- +goose Down
DROP TABLE IF EXISTS chats;