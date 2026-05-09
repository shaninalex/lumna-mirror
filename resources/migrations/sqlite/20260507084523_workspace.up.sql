CREATE TABLE workspaces
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    title       text     NOT NULL,
    active      numeric,
    owner_email text     NULL,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at  datetime NULL
);

ALTER TABLE invitations
    ADD COLUMN workspace_id INTEGER REFERENCES workspaces (id);
CREATE INDEX idx_invitations_workspace_id ON invitations (workspace_id);

ALTER TABLE projects
    ADD COLUMN workspace_id INTEGER REFERENCES workspaces (id);
CREATE INDEX idx_projects_workspace_id ON projects (workspace_id);

ALTER TABLE boards
    ADD COLUMN workspace_id INTEGER REFERENCES workspaces (id);
CREATE INDEX idx_boards_workspace_id ON boards (workspace_id);

ALTER TABLE columns
    ADD COLUMN workspace_id INTEGER REFERENCES workspaces (id);
CREATE INDEX idx_columns_workspace_id ON columns (workspace_id);

ALTER TABLE tasks
    ADD COLUMN workspace_id INTEGER REFERENCES workspaces (id);
CREATE INDEX idx_tasks_workspace_id ON tasks (workspace_id);
