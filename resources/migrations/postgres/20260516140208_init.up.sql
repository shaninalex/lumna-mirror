--- identities definition
CREATE TABLE identities
(
    id         SERIAL PRIMARY KEY,
    full_name  VARCHAR NOT NULL,
    email      VARCHAR NOT NULL,
    active     NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT uni_identities_email UNIQUE (email)
);

-- Credentials
CREATE TABLE credentials
(
    id               SERIAL PRIMARY KEY,
    identity_id      INTEGER NOT NULL,
    provider         VARCHAR NOT NULL,
    provider_user_id VARCHAR,
    email            VARCHAR,
    password_hash    VARCHAR,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_credentials_identity FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX idx_provider_email ON credentials (provider, email);
CREATE UNIQUE INDEX idx_provider_user ON credentials (provider, provider_user_id);
CREATE INDEX idx_credentials_identity_id ON credentials (identity_id);

-- refresh_tokens definition
CREATE TABLE refresh_tokens
(
    id          SERIAL PRIMARY KEY,
    identity_id INTEGER NOT NULL,
    hash        VARCHAR NOT NULL,
    client_id   VARCHAR,
    scopes      VARCHAR,
    expires_at  TIMESTAMP NOT NULL,
    revoked     BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_refresh_tokens_identity FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens (expires_at);
CREATE UNIQUE INDEX idx_refresh_tokens_hash ON refresh_tokens (hash);
CREATE INDEX idx_refresh_tokens_identity_id ON refresh_tokens (identity_id);

-- invitations definition
CREATE TABLE invitations
(
    id          SERIAL PRIMARY KEY,
    email       VARCHAR,
    token_hash  VARCHAR,
    state       VARCHAR,
    role        VARCHAR,
    created_at  TIMESTAMP NOT NULL,
    valid_until TIMESTAMP NOT NULL,
    meta        TEXT        NULL,

    CONSTRAINT uni_invitations_email UNIQUE (email)
);
CREATE UNIQUE INDEX idx_invitations_token_hash ON invitations (token_hash);

CREATE TABLE workspaces
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR NOT NULL,
    slug        VARCHAR NOT NULL,
    active      NUMERIC,
    owner_email VARCHAR NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NULL
);

CREATE TABLE projects
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR NOT NULL,
    workspace_id INTEGER REFERENCES workspaces (id),
    created_at   TIMESTAMP,
    updated_at   TIMESTAMP
);
CREATE TABLE lists
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR NOT NULL,
    project_id INTEGER NOT NULL,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    workspace_id INTEGER REFERENCES workspaces (id),
    CONSTRAINT fk_projects_lists FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_lists_project_id ON lists (project_id);

CREATE TABLE statuses
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR NOT NULL,
    "order"    INTEGER,
    list_id    INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    created_at  TIMESTAMP,
    updated_at   TIMESTAMP,
    CONSTRAINT fk_status_project FOREIGN KEY (project_id) REFERENCES projects (id),
    CONSTRAINT fk_status_list FOREIGN KEY (list_id) REFERENCES lists (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_columns_project_id ON statuses (project_id);
CREATE INDEX idx_columns_board_id ON statuses (list_id); -- Assuming "board_id" is a typo, changed to "list_id" to match table definition

CREATE TABLE tasks
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR NOT NULL,
    "order"    INTEGER,
    done       BOOLEAN,
    body       VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    status_id  INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    list_id    INTEGER NOT NULL,
    workspace_id INTEGER REFERENCES workspaces (id),
    CONSTRAINT fk_tasks_project FOREIGN KEY (project_id) REFERENCES projects (id),
    CONSTRAINT fk_tasks_list FOREIGN KEY (list_id) REFERENCES lists (id),
    CONSTRAINT fk_tasks_status FOREIGN KEY (status_id) REFERENCES statuses (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_tasks_list_id ON tasks (list_id);
CREATE INDEX idx_tasks_project_id ON tasks (project_id);
CREATE INDEX idx_tasks_status_id ON tasks (status_id);

CREATE TABLE activity_logs
(
    id          SERIAL PRIMARY KEY,
    summary     VARCHAR NOT NULL,
    identity_id INTEGER,
    entity_id   INTEGER NOT NULL,
    entity_type VARCHAR NOT NULL,
    action      VARCHAR NOT NULL,
    created_at  TIMESTAMP,
    CONSTRAINT fk_activity_logs_identity FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_activity_logs_action ON activity_logs (action);
CREATE INDEX idx_activity_logs_entity_type ON activity_logs (entity_type);
CREATE INDEX idx_activity_logs_entity_id ON activity_logs (entity_id);
CREATE INDEX idx_activity_logs_identity_id ON activity_logs (identity_id);