--- identities definition
CREATE TABLE identities
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    full_name  text NOT NULL,
    email      text NOT NULL,
    active     numeric,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime,
    CONSTRAINT uni_identities_email UNIQUE (email)
);

-- Credentials
CREATE TABLE credentials
(
    id               integer PRIMARY KEY AUTOINCREMENT,
    identity_id      integer NOT NULL,
    provider         text    NOT NULL,
    provider_user_id text,
    email            text,
    password_hash    text,
    created_at       datetime DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_credentials_identity FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX idx_provider_email ON credentials (provider, email);
CREATE UNIQUE INDEX idx_provider_user ON credentials (provider, provider_user_id);
CREATE INDEX idx_credentials_identity_id ON credentials (identity_id);

-- refresh_tokens definition
CREATE TABLE refresh_tokens
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    identity_id integer  NOT NULL,
    hash        text     NOT NULL,
    client_id   text,
    scopes      text,
    expires_at  datetime NOT NULL,
    revoked     numeric  NOT NULL DEFAULT false,
    created_at  datetime          DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_refresh_tokens_identity FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens (expires_at);
CREATE UNIQUE INDEX idx_refresh_tokens_hash ON refresh_tokens (hash);
CREATE INDEX idx_refresh_tokens_identity_id ON refresh_tokens (identity_id);

-- invitations definition
CREATE TABLE invitations
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    email       text,
    token_hash  text,
    state       text,
    role        text,
    created_at  datetime NOT NULL,
    valid_until datetime NOT NULL,
    meta        text     NULL,

    CONSTRAINT uni_invitations_email UNIQUE (email)
);
CREATE UNIQUE INDEX idx_invitations_token_hash ON invitations (token_hash);

CREATE TABLE workspaces
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    title       text     NOT NULL,
    active      numeric,
    owner_email text     NULL,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at  datetime NULL
);

CREATE TABLE projects
(
    id           integer PRIMARY KEY AUTOINCREMENT,
    title        text    NOT NULL,
    workspace_id INTEGER REFERENCES workspaces (id) ON DELETE SET NULL,
    created_at   datetime,
    updated_at   datetime
);

CREATE TABLE lists
(
    id           integer PRIMARY KEY AUTOINCREMENT,
    title        text    NOT NULL,
    project_id   integer NOT NULL,
    created_at   datetime,
    updated_at   datetime,
    CONSTRAINT fk_projects_lists FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_lists_project_id ON lists (project_id);

CREATE TABLE statuses
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    `order`    integer,
    list_id    integer NOT NULL REFERENCES lists (id) ON DELETE CASCADE ON UPDATE CASCADE,
    project_id integer NOT NULL REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at datetime,
    updated_at datetime
);
CREATE INDEX idx_columns_project_id ON statuses (project_id);
CREATE INDEX idx_columns_list_id ON statuses (list_id);

CREATE TABLE tasks
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    `order`    integer,
    done       numeric,
    body       text,
    created_at datetime,
    updated_at datetime,
    status_id  integer NOT NULL,
    project_id integer NOT NULL,

    CONSTRAINT fk_tasks_project FOREIGN KEY (project_id) REFERENCES projects (id),
    CONSTRAINT fk_tasks_status FOREIGN KEY (status_id) REFERENCES statuses (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_tasks_project_id ON tasks (project_id);
CREATE INDEX idx_tasks_status_id ON tasks (status_id);

CREATE TABLE activity_logs
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    summary     text    NOT NULL,
    identity_id integer,
    entity_id   integer NOT NULL,
    entity_type text    NOT NULL,
    action      text    NOT NULL,
    created_at  datetime,
    CONSTRAINT fk_activity_logs_identity FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);
