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

CREATE TABLE user_auth_history
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    event_type  text NOT NULL,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP
);

-----------------------------
-- SYSTEM
-----------------------------

CREATE TABLE emails
(
    id          integer PRIMARY KEY AUTOINCREMENT,
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

/*

Core entities:
    Workspace
    Project
    Board
    Column
    Task

Relationships:
    Project -> Workspace
    Board -> Project
    Column -> Board
    Taks -> Column

Events/History:
    TaskHistory

Value objects/attributes:
    Color
    Tag/Label
    Badge
    (badge is not a tag/label. Badge can have logic in it, while label is just a lable)
*/

CREATE TABLE projects
(
    id           integer PRIMARY KEY AUTOINCREMENT,
    title        text    NOT NULL,
    workspace_id integer NOT NULL,
    owner_id     integer NULL,
    meta         text    NULL,
    created_at   datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at   datetime,

    FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE,
    FOREIGN KEY (owner_id) REFERENCES identities (id) ON DELETE SET NULL
);

CREATE TABLE boards
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    project_id integer NOT NULL,
    meta       text,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);


CREATE TABLE columns
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    board_id  integer NOT NULL,
    meta       varchar null,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime,

    FOREIGN KEY (board_id) REFERENCES boards(id) ON DELETE CASCADE
);

CREATE TABLE tasks
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    title       text NOT NULL,
    body        text,
    completed   boolean DEFAULT false,
    meta        text,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME
);

CREATE TABLE board_tasks
(
    board_id  integer NOT NULL,
    task_id   integer NOT NULL,
    column_id integer NULL,
    position  real NOT NULL,

    PRIMARY KEY (board_id, task_id),

    FOREIGN KEY (board_id)
        REFERENCES boards(id)
        ON DELETE CASCADE,

    FOREIGN KEY (task_id)
        REFERENCES tasks(id)
        ON DELETE CASCADE,

    FOREIGN KEY (board_id, column_id)
        REFERENCES columns(board_id, id)
        ON DELETE SET NULL
);

CREATE TABLE task_events
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    identity_id integer NULL,
    event_type  text NOT NULL,
    value_from  text NULL,
    value_to    text NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE SET NULL
);
