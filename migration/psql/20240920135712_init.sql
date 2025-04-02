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

CREATE TABLE IF NOT EXISTS account_events (
    id VARCHAR(50) PRIMARY KEY,
    account_id VARCHAR(50) references accounts (id) NOT NULL,
    event_type VARCHAR(124) NOT NULL,
    event_data JSONB,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS event_send (
    events_id VARCHAR(50) PRIMARY KEY references account_events (id) NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
