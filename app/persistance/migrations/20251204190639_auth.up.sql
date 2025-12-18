CREATE TABLE users_tokens
(
    id                 INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id            INTEGER NOT NULL,
    device             VARCHAR,
    refresh_token      VARCHAR NOT NULL,
    refresh_expires_at DATETIME,
    revoked            BOOLEAN  DEFAULT 0,
    revoked_at         DATETIME,
    created_at         DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- CREATE UNIQUE INDEX idx_users_tokens_refresh_token ON users_tokens (refresh_token);