CREATE TABLE sprints
(
    id          SERIAL PRIMARY KEY,
    title       TEXT      NOT NULL,
    description TEXT,
    done        INTEGER   NOT NULL DEFAULT 0,
    project_id  INTEGER   NOT NULL,
    started_at  TIMESTAMP,
    finished_at TIMESTAMP,

    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,

    FOREIGN KEY (project_id) REFERENCES projects (id)
);