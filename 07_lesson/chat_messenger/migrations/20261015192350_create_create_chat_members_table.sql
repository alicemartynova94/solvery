-- +goose Up
CREATE TABLE chat_members
(
    chat_id    UUID NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    session_id UUID NOT NULL REFERENCES sessions (id),
    PRIMARY KEY (chat_id, session_id)
);

-- +goose Down
DROP TABLE IF EXISTS chat_members;
