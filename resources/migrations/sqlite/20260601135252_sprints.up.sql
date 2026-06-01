CREATE TABLE sprints
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT     NOT NULL,
    description TEXT,
    done        INTEGER  NOT NULL DEFAULT 0,
    project_id  INTEGER  NOT NULL,
    started_at  DATETIME,
    finished_at DATETIME,

    created_at  DATETIME NOT NULL,
    updated_at  DATETIME NOT NULL,

    FOREIGN KEY (project_id) REFERENCES projects (id)
);