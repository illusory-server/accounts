-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(50) PRIMARY KEY,
    nickname VARCHAR(128) NOT NULL UNIQUE,
    first_name VARCHAR(128),
    last_name VARCHAR(128),
    email VARCHAR(255) NOT NULL UNIQUE,
    role VARCHAR(128) NOT NULL,
    avatar_link VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
