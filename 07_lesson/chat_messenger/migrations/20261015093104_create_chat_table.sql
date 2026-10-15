-- +goose Up
CREATE TABLE chats
(
    id           UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    chat_name    VARCHAR(100) NOT NULL,
    expires_at   TIMESTAMPTZ,
    read_only    BOOLEAN      NOT NULL DEFAULT FALSE,
    max_messages INTEGER      NOT NULL DEFAULT 0,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS chats;