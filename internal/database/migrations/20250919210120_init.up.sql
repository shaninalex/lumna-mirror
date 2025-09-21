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


CREATE TABLE users_tokens
(
    id         TEXT PRIMARY KEY NOT NULL,
    user_id    TEXT             NOT NULL,
    -- token data, permissions, other things
    claims     TEXT             NOT NULL,
    device     TEXT,
    expires_at DATETIME         NOT NULL,
    created_at DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);