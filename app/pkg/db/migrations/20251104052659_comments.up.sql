CREATE TABLE comments
(
    id          INTEGER     PRIMARY KEY AUTOINCREMENT,
    entity_id   INTEGER     NOT NULL,
    entity_type VARCHAR     NOT NULL,
    user_id     INTEGER     NOT NULL,
    content     VARCHAR     NOT NULL,
    created_at  DATETIME    DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
