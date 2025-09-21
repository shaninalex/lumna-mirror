CREATE TABLE projects
(
    id              TEXT PRIMARY KEY NOT NULL,
    user_id         TEXT             NOT NULL,
    organization_id TEXT             NULL,
    title           TEXT             NOT NULL,
    code            TEXT             NOT NULL,
    created_at      DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
