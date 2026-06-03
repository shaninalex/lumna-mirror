-----------------------------
-- IDENTITIES AND CREDENTIALS
-----------------------------

CREATE TABLE identities
(
    id         SERIAL PRIMARY KEY,
    full_name  VARCHAR NOT NULL,
    email      VARCHAR NOT NULL,
    active     BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,

    CONSTRAINT uni_identities_email UNIQUE (email)
);

CREATE TABLE credentials
(
    id               SERIAL PRIMARY KEY,
    identity_id      INTEGER NOT NULL,
    provider         VARCHAR NOT NULL,
    provider_user_id VARCHAR,
    email            VARCHAR,
    password_hash    VARCHAR,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE refresh_tokens
(
    id          SERIAL PRIMARY KEY,
    identity_id INTEGER NOT NULL,
    hash        VARCHAR NOT NULL,
    client_id   VARCHAR,
    scopes      VARCHAR,
    expires_at  TIMESTAMP NOT NULL,
    revoked     BOOLEAN NOT NULL DEFAULT false,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-----------------------------
-- SYSTEM
-----------------------------

CREATE TABLE activity_logs
(
    id          SERIAL PRIMARY KEY,
    summary     VARCHAR NOT NULL,
    identity_id INTEGER,
    entity_id   INTEGER NOT NULL,
    entity_type VARCHAR NOT NULL,
    action      VARCHAR NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE emails
(
    id          SERIAL PRIMARY KEY,
    from_email  VARCHAR,
    to_email    VARCHAR,
    subject     VARCHAR,
    body        VARCHAR,
    format      VARCHAR,
    sender_name VARCHAR,
    headers     VARCHAR,
    status      VARCHAR
        CHECK (status IN ('pending', 'running', 'success', 'repeat', 'error', 'skipped'))
                          DEFAULT 'pending',
    cc          VARCHAR,
    bcc         VARCHAR,
    reply_to    VARCHAR,
    sent_at     TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    meta        TEXT NULL
);

-----------------------------
-- WORKSPACES AND USER MANAGEMENT
-----------------------------

CREATE TABLE invitations
(
    id          SERIAL PRIMARY KEY,
    email       VARCHAR,
    token_hash  VARCHAR,
    state       VARCHAR,
    role        VARCHAR,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valid_until TIMESTAMP NOT NULL,
    meta        TEXT NULL,

    CONSTRAINT uni_invitations_email UNIQUE (email)
);

CREATE TABLE workspaces
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR NOT NULL,
    active      BOOLEAN DEFAULT true,
    owner_email VARCHAR NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NULL
);

-----------------------------
-- WORK
-----------------------------

CREATE TABLE projects
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR NOT NULL,
    workspace_id INTEGER NOT NULL,
    owner_id     INTEGER NULL,
    created_at   TIMESTAMP,
    updated_at   TIMESTAMP,

    FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE,
    FOREIGN KEY (owner_id) REFERENCES identities (id) ON DELETE SET NULL
);

CREATE TABLE sprints
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR NOT NULL,
    description VARCHAR,
    done        BOOLEAN NOT NULL DEFAULT false,
    project_id  INTEGER NOT NULL,
    started_at  TIMESTAMP NULL,
    finished_at TIMESTAMP NULL,

    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NULL,

    FOREIGN KEY (project_id) REFERENCES projects (id)
);

CREATE TABLE lists
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR NOT NULL,
    project_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE statuses
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR NOT NULL,
    "order"    INTEGER NULL,
--     list_id    INTEGER NOT NULL REFERENCES lists (id) ON DELETE CASCADE ON UPDATE CASCADE,
    project_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE tasks
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR NOT NULL,
    "order"    INTEGER,
    done       BOOLEAN DEFAULT false,
    body       VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    status_id  INTEGER NOT NULL,
    project_id INTEGER NOT NULL,

    FOREIGN KEY (status_id) REFERENCES statuses (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);