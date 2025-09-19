CREATE TABLE users
(
    id              TEXT PRIMARY KEY,
    organization_id TEXT NOT NULL,
    email           TEXT NOT NULL,
    settings        TEXT,
    active          BOOLEAN   DEFAULT false,
    state           TEXT      DEFAULT 'pending',
    code            TEXT,
    password_hash   TEXT NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);