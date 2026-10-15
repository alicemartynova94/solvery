-- +goose Up
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sessions
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPZ
);

-- +goose Down
DROP TABLE IF EXISTS sessions;
