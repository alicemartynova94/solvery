-- +goose Up
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sessions
(
    id         UUID PRIMARY KEY    DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- +goose Down
DROP TABLE IF EXISTS sessions;
DROP EXTENSION IF EXISTS "uuid-ossp";
