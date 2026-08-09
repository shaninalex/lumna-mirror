-----------------------------
-- IDENTITIES AND CREDENTIALS
-----------------------------

CREATE TABLE identities
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    full_name  text     NOT NULL,
    email      text     NOT NULL,
    active     numeric,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NULL,

    CONSTRAINT uni_identities_email UNIQUE (email)
);

CREATE TABLE credentials
(
    id               integer PRIMARY KEY AUTOINCREMENT,
    identity_id      integer NOT NULL,
    provider         text    NOT NULL,
    provider_user_id text,
    email            text,
    password_hash    text,
    created_at       datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

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
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-----------------------------
-- SYSTEM
-----------------------------

CREATE TABLE emails
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
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
    sent_at     DATETIME,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    meta        text NULL
);

-----------------------------
-- WORKSPACES AND USER MANAGEMENT
-----------------------------

CREATE TABLE invitations
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    email       text,
    token_hash  text,
    state       text,
    role        text,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP,
    valid_until datetime NOT NULL,
    meta        text     NULL,

    CONSTRAINT uni_invitations_email UNIQUE (email)
);

CREATE TABLE workspaces
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    title       text     NOT NULL,
    active      numeric,
    owner_email text     NULL,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at  datetime NULL
);

-----------------------------
-- WORK
-----------------------------

CREATE TABLE projects
(
    id           integer PRIMARY KEY AUTOINCREMENT,
    title        text    NOT NULL,
    key          varchar not null unique,
    workspace_id INTEGER NOT NULL,
    owner_id     INTEGER NULL,
    meta         text    NULL,
    created_at   datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at   datetime,

    FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE,
    FOREIGN KEY (owner_id) REFERENCES identities (id) ON DELETE SET NULL
);


CREATE TABLE lists -- // boards
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    project_id integer NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE statuses -- // columns
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    `order`    integer NULL,
    project_id integer NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE tasks
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    code       varchar NOT NULL UNIQUE,
    `order`    integer,
    done       numeric  DEFAULT 0,
    body       text,
    status_id  integer NULL,
    project_id integer NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime,

    FOREIGN KEY (status_id) REFERENCES statuses (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
);