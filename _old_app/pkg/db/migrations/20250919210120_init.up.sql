CREATE TABLE users
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    email         TEXT NOT NULL,
    settings      TEXT,
    active        BOOLEAN   DEFAULT false,
    state         VARCHAR   DEFAULT 'pending',
    code          VARCHAR,
    password_hash TEXT NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT users_unique_email_code UNIQUE (email, code)
);
CREATE UNIQUE INDEX idx_users_code ON users (code);

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

CREATE UNIQUE INDEX idx_users_tokens_refresh_token ON users_tokens (refresh_token);

CREATE TABLE projects
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    INTEGER  NOT NULL,
    title      VARCHAR  NOT NULL,
    code       VARCHAR  NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,

    CONSTRAINT projects_unique_title UNIQUE (title),
    CONSTRAINT projects_unique_code UNIQUE (code)
);

CREATE TABLE statuses
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id INTEGER NOT NULL,
    title      VARCHAR NOT NULL,
    list_index INT     default 0,
    config     VARCHAR,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE tasks
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id     INTEGER  NOT NULL,
    project_id  INTEGER  NOT NULL,
    status_id   INTEGER  NOT NULL,
    title       VARCHAR  NOT NULL,
    code        VARCHAR  NOT NULL,
    completed   BOOLEAN           DEFAULT false,
    description VARCHAR  NULL,
    list_index  FLOAT             default 0,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    FOREIGN KEY (status_id) REFERENCES statuses (id) ON DELETE CASCADE,

    CONSTRAINT tasks_unique_code UNIQUE (code)
);
