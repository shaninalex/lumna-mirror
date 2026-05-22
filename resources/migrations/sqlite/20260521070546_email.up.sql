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
        CHECK(status IN ('pending', 'running', 'success', 'repeat', 'error', 'skipped'))
        NOT NULL DEFAULT 'pending',

    cc          VARCHAR,
    bcc         VARCHAR,
    reply_to    VARCHAR,
    sent_at     DATETIME,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    meta        text        NULL
);