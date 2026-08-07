-----------------------------
-- IDENTITIES AND CREDENTIALS
-----------------------------

CREATE TABLE identities
(
    id         SERIAL PRIMARY KEY,
    full_name  TEXT     NOT NULL,
    email      TEXT     NOT NULL,
    active     NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,

    CONSTRAINT uni_identities_email UNIQUE (email)
);

CREATE TABLE credentials
(
    id               SERIAL PRIMARY KEY,
    identity_id      INTEGER NOT NULL,
    provider         TEXT    NOT NULL,
    provider_user_id TEXT,
    email            TEXT,
    password_hash    TEXT,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE refresh_tokens
(
    id          SERIAL PRIMARY KEY,
    identity_id INTEGER  NOT NULL,
    hash        TEXT     NOT NULL,
    client_id   TEXT,
    scopes      TEXT,
    expires_at  TIMESTAMP NOT NULL,
    revoked     NUMERIC  NOT NULL DEFAULT false,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-----------------------------
-- SYSTEM
-----------------------------

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
    email       TEXT,
    token_hash  TEXT,
    state       TEXT,
    role        TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valid_until TIMESTAMP NOT NULL,
    meta        TEXT     NULL,

    CONSTRAINT uni_invitations_email UNIQUE (email)
);

CREATE TABLE workspaces
(
    id          SERIAL PRIMARY KEY,
    title       TEXT     NOT NULL,
    active      NUMERIC,
    owner_email TEXT     NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NULL
);

-----------------------------
-- WORK
-----------------------------

CREATE TABLE projects
(
    id           SERIAL PRIMARY KEY,
    title        TEXT    NOT NULL,
    workspace_id INTEGER NOT NULL,
    owner_id     INTEGER NULL,
    key          VARCHAR NOT NULL UNIQUE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP,
    meta         TEXT    NULL,

    FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE,
    FOREIGN KEY (owner_id) REFERENCES identities (id) ON DELETE SET NULL
);


CREATE TABLE lists -- // boards
(
    id         SERIAL PRIMARY KEY,
    title      TEXT    NOT NULL,
    project_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE statuses -- // columns
(
    id         SERIAL PRIMARY KEY,
    title      TEXT    NOT NULL,
    "order"    INTEGER NULL,
    project_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE tasks
(
    id         SERIAL PRIMARY KEY,
    title      TEXT    NOT NULL,
    code       VARCHAR NOT NULL UNIQUE,
    "order"    INTEGER,
    done       NUMERIC  DEFAULT 0,
    body       TEXT,
    status_id  INTEGER NULL,
    project_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (status_id) REFERENCES statuses (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);