-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    nickname TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    role TEXT NOT NULL,
    avatar_link TEXT,
    password TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    version BIGINT DEFAULT 0,

    CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    CONSTRAINT valid_role CHECK (role IN ('USER', 'ADMIN', 'SUPER_ADMIN')),
    CONSTRAINT nickname_length CHECK (LENGTH(nickname) BETWEEN 2 AND 128),

    UNIQUE (nickname),
    UNIQUE (email)
);

CREATE INDEX idx_accounts_created ON accounts (created_at);

CREATE TABLE IF NOT EXISTS account_events (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id),
    event_type TEXT NOT NULL,
    event_data JSONB NOT NULL CHECK (jsonb_typeof(event_data) = 'object'),
    timestamp TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    aggregate_version BIGINT NOT NULL,
    metadata JSONB,

    CONSTRAINT valid_event_data CHECK (event_data IS NOT NULL)
);

-- CREATE TABLE account_events_2023_11 PARTITION OF account_events
--     FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

-- CREATE TABLE account_events_default PARTITION OF account_events DEFAULT;

CREATE INDEX idx_account_events_account ON account_events (account_id);
CREATE INDEX idx_account_events_type ON account_events (event_type);
CREATE INDEX idx_account_events_version ON account_events (account_id, aggregate_version);

ALTER TABLE account_events SET (
    autovacuum_vacuum_scale_factor = 0,
    autovacuum_vacuum_threshold = 10000
);

CREATE TABLE IF NOT EXISTS account_outbox (
    id TEXT NOT NULL,
    event_id TEXT NOT NULL REFERENCES account_events(id),
    event_type TEXT NOT NULL CHECK (event_type <> ''),
    payload JSONB NOT NULL CHECK (jsonb_typeof(payload) = 'object'),
    status TEXT NOT NULL DEFAULT 'pending'
      CHECK (status IN ('pending', 'processing', 'processed', 'failed')),
    retry_count INTEGER NOT NULL DEFAULT 0 CHECK (retry_count >= 0),
    error_message TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    scheduled_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ,

    PRIMARY KEY (id, status, created_at),

    CONSTRAINT unique_event_id UNIQUE (event_id, status, created_at),

    priority SMALLINT NOT NULL DEFAULT 100 CHECK (priority BETWEEN 0 AND 255),

    partition_key TEXT NOT NULL DEFAULT TO_CHAR(NOW(), 'YYYY_MM')
) PARTITION BY LIST (status);

CREATE TABLE outbox_pending PARTITION OF account_outbox
    FOR VALUES IN ('pending');
CREATE TABLE outbox_processing PARTITION OF account_outbox
    FOR VALUES IN ('processing');
CREATE TABLE outbox_processed PARTITION OF account_outbox
    FOR VALUES IN ('processed') PARTITION BY RANGE (created_at);
CREATE TABLE outbox_failed PARTITION OF account_outbox
    FOR VALUES IN ('failed');

-- +goose StatementEnd

INSERT INTO accounts (
    id,
    first_name,
    last_name,
    email,
    role,
    nickname,
    password,
    updated_at,
    created_at
) VALUES (
    'a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8', -- UUID
    'Иван',                                  -- first_name
    'Петров',                               -- last_name
    'ivan.petrov@example.com',              -- email
    'USER',                                 -- role
    'ivan_the_terrible',                    -- nickname
    '$2a$10$xJwL5v5zJZUfQ7zvWbUYr.XLz6tB5Gd9V6bYwL0aNcF1kZw2sYbW', -- хеш пароля
    NOW(),                                  -- updated_at
    NOW()                                   -- created_at
);

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
