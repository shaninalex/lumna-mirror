PRAGMA foreign_keys = ON;

-------------------------
-- USERS
-------------------------
CREATE TABLE users
(
    id     INTEGER PRIMARY KEY,
    email  TEXT    NOT NULL UNIQUE,
    active BOOLEAN NOT NULL DEFAULT 1
);

-------------------------
-- PROJECT
-------------------------
CREATE TABLE projects
(
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

-------------------------
-- BOARDS (child of project)
-------------------------
CREATE TABLE boards
(
    id         INTEGER PRIMARY KEY,
    project_id INTEGER NOT NULL,
    settings   TEXT,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

-------------------------
-- LISTS (child of board)
-------------------------
CREATE TABLE lists
(
    id       INTEGER PRIMARY KEY,
    board_id INTEGER NOT NULL,
    name     TEXT    NOT NULL,
    "order"  INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (board_id) REFERENCES boards (id) ON DELETE CASCADE
);

-------------------------
-- TASKS (child of board, linked by ListId)
-------------------------
CREATE TABLE tasks
(
    id         INTEGER PRIMARY KEY,
    board_id   INTEGER  NOT NULL,
    list_id    INTEGER  NOT NULL,
    name       TEXT     NOT NULL,
    done       BOOLEAN  NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (board_id) REFERENCES boards (id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES lists (id) ON DELETE SET NULL
);

-- Optional index
CREATE INDEX idx_tasks_board ON tasks (board_id);
CREATE INDEX idx_tasks_list ON tasks (list_id);

-------------------------
-- CALENDARS (child of project)
-------------------------
CREATE TABLE calendars
(
    id         INTEGER PRIMARY KEY,
    project_id INTEGER NOT NULL,
    settings   TEXT,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

-------------------------
-- CALENDAR ENTRIES
-------------------------
CREATE TABLE calendar_entries
(
    id          INTEGER PRIMARY KEY,
    calendar_id INTEGER  NOT NULL,
    date        DATETIME NOT NULL,
    name        TEXT     NOT NULL,
    FOREIGN KEY (calendar_id) REFERENCES calendars (id) ON DELETE CASCADE
);

-------------------------
-- DOCUMENTS & FOLDERS (tree)
-------------------------

CREATE TABLE folders
(
    id         INTEGER PRIMARY KEY,
    project_id INTEGER NOT NULL,
    name       TEXT    NOT NULL,
    parent_id  INTEGER,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES folders (id) ON DELETE CASCADE
);

CREATE TABLE documents
(
    id        INTEGER PRIMARY KEY,
    name      TEXT    NOT NULL,
    folder_id INTEGER NOT NULL,
    FOREIGN KEY (folder_id) REFERENCES folders (id) ON DELETE CASCADE
);

-------------------------
-- COMMENTS (generic entity comments)
-------------------------
CREATE TABLE comments
(
    id          INTEGER PRIMARY KEY,
    date        DATETIME NOT NULL,
    content     TEXT     NOT NULL,
    entity_type TEXT     NOT NULL, -- e.g., "task", "document", etc.
    entity_id   INTEGER  NOT NULL,
    author_id   INTEGER  NOT NULL,
    created_at  DATETIME NOT NULL,
    updated_at  DATETIME NOT NULL,

    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE SET NULL
);

CREATE INDEX idx_comments_entity ON comments (entity_type, entity_id);
