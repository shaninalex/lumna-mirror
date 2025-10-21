CREATE TABLE badge
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id INTEGER  NOT NULL,
    title      VARCHAR  NOT NULL,
    config     VARCHAR,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE badges_tasks
(
    badge_id INTEGER NOT NULL,
    task_id    INTEGER NOT NULL,

    FOREIGN KEY (badge_id) REFERENCES badge (id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE,

    CONSTRAINT badges_tasks_unique UNIQUE (badge_id, task_id)
);