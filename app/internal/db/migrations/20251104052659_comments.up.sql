CREATE TABLE comments
(
    id          INTEGER     PRIMARY KEY AUTOINCREMENT,
    task_id     INTEGER     NOT NULL,
    user_id     INTEGER     NOT NULL,
    content     VARCHAR     NOT NULL,
    created_at  DATETIME    DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

