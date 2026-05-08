--- IDENTITY ---
-- identities definition
CREATE TABLE "identities"
(
    "id"         SERIAL PRIMARY KEY,
    "full_name"  TEXT NOT NULL,
    "email"      TEXT NOT NULL,
    "active"     NUMERIC,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    CONSTRAINT "uni_identities_email" UNIQUE ("email")
);

-- Credentials
CREATE TABLE "credentials"
(
    "identity_id"      INTEGER   NOT NULL,
    "provider"         TEXT      NOT NULL,
    "provider_user_id" TEXT,
    "email"            TEXT,
    "password_hash"    TEXT,
    "created_at"       TIMESTAMP NOT NULL,
    CONSTRAINT "fk_credentials_identity" FOREIGN KEY ("identity_id") REFERENCES "identities" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX "idx_provider_email" ON "credentials" ("provider", "email");
CREATE UNIQUE INDEX "idx_provider_user" ON "credentials" ("provider", "provider_user_id");
CREATE INDEX "idx_credentials_identity_id" ON "credentials" ("identity_id");

-- refresh_tokens definition
CREATE TABLE "refresh_tokens"
(
    "id"          SERIAL PRIMARY KEY,
    "identity_id" INTEGER   NOT NULL,
    "hash"        TEXT      NOT NULL,
    "client_id"   TEXT,
    "scopes"      TEXT,
    "expires_at"  TIMESTAMP NOT NULL,
    "revoked"     BOOLEAN   NOT NULL DEFAULT FALSE,
    "created_at"  TIMESTAMP NOT NULL,
    CONSTRAINT "fk_refresh_tokens_identity" FOREIGN KEY ("identity_id") REFERENCES "identities" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX "idx_refresh_tokens_expires_at" ON "refresh_tokens" ("expires_at");
CREATE UNIQUE INDEX "idx_refresh_tokens_hash" ON "refresh_tokens" ("hash");
CREATE INDEX "idx_refresh_tokens_identity_id" ON "refresh_tokens" ("identity_id");

-- invitations definition
CREATE TABLE "invitations"
(
    "id"          SERIAL PRIMARY KEY,
    "email"       TEXT,
    "token_hash"  TEXT,
    "state"       TEXT,
    "role"        TEXT,
    "created_at"  TIMESTAMP NOT NULL,
    "valid_until" TIMESTAMP NOT NULL,
    CONSTRAINT "uni_invitations_email" UNIQUE ("email")
);
CREATE UNIQUE INDEX "idx_invitations_token_hash" ON "invitations" ("token_hash");

--- PROJECT ---
-- projects definition
CREATE TABLE "projects"
(
    "id"         SERIAL PRIMARY KEY,
    "title"      TEXT NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

-- Boards table
CREATE TABLE "boards"
(
    "id"         SERIAL PRIMARY KEY,
    "title"      TEXT    NOT NULL,
    "project_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    CONSTRAINT "fk_projects_boards" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX "idx_boards_project_id" ON "boards" ("project_id");

-- Columns
CREATE TABLE "columns"
(
    "id"         SERIAL PRIMARY KEY,
    "title"      TEXT    NOT NULL,
    "order"      INTEGER,
    "board_id"   INTEGER NOT NULL,
    "project_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    CONSTRAINT "fk_columns_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id"),
    CONSTRAINT "fk_boards_columns" FOREIGN KEY ("board_id") REFERENCES "boards" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX "idx_columns_project_id" ON "columns" ("project_id");
CREATE INDEX "idx_columns_board_id" ON "columns" ("board_id");

-- tasks definition
CREATE TABLE "tasks"
(
    "id"         SERIAL PRIMARY KEY,
    "title"      TEXT    NOT NULL,
    "order"      INTEGER,
    "done"       NUMERIC,
    "body"       TEXT,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "column_id"  INTEGER NOT NULL,
    "project_id" INTEGER NOT NULL,
    "board_id"   INTEGER NOT NULL,
    CONSTRAINT "fk_tasks_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id"),
    CONSTRAINT "fk_tasks_board" FOREIGN KEY ("board_id") REFERENCES "boards" ("id"),
    CONSTRAINT "fk_columns_tasks" FOREIGN KEY ("column_id") REFERENCES "columns" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX "idx_tasks_board_id" ON "tasks" ("board_id");
CREATE INDEX "idx_tasks_project_id" ON "tasks" ("project_id");
CREATE INDEX "idx_tasks_column_id" ON "tasks" ("column_id");

--- OTHER ---
-- activity_logs
CREATE TABLE "activity_logs"
(
    "id"          SERIAL PRIMARY KEY,
    "summary"     TEXT    NOT NULL,
    "identity_id" INTEGER NOT NULL,
    "entity_id"   INTEGER NOT NULL,
    "entity_type" TEXT    NOT NULL,
    "action"      TEXT    NOT NULL,
    "created_at"  TIMESTAMP,
    CONSTRAINT "fk_activity_logs_identity" FOREIGN KEY ("identity_id") REFERENCES "identities" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX "idx_activity_logs_action" ON "activity_logs" ("action");
CREATE INDEX "idx_activity_logs_entity_type" ON "activity_logs" ("entity_type");
CREATE INDEX "idx_activity_logs_entity_id" ON "activity_logs" ("entity_id");
CREATE INDEX "idx_activity_logs_identity_id" ON "activity_logs" ("identity_id");

-- jobs definition
CREATE TABLE "jobs"
(
    "id"           SERIAL PRIMARY KEY,
    "type"         TEXT,
    "payload"      TEXT,
    "status"       TEXT,
    "attempts"     INTEGER,
    "available_at" TIMESTAMP,
    "locked_at"    TIMESTAMP,
    "locked_by"    TEXT,
    "created_at"   TIMESTAMP
);
