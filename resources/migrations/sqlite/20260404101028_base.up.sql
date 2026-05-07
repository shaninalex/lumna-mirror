--- IDENTITY ---
--- identities definition
CREATE TABLE `identities` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `full_name` text NOT NULL,
    `email` text NOT NULL,
    `active` numeric,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime,
    CONSTRAINT `uni_identities_email` UNIQUE (`email`)
);

-- Creadentials
CREATE TABLE `credentials` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `identity_id` integer NOT NULL,
    `provider` text NOT NULL,
    `provider_user_id` text,
    `email` text,
    `password_hash` text,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_credentials_identity` FOREIGN KEY (`identity_id`) REFERENCES `identities`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX `idx_provider_email` ON `credentials`(`provider`, `email`);
CREATE UNIQUE INDEX `idx_provider_user` ON `credentials`(`provider`, `provider_user_id`);
CREATE INDEX `idx_credentials_identity_id` ON `credentials`(`identity_id`);

-- refresh_tokens definition
CREATE TABLE `refresh_tokens` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `identity_id` integer NOT NULL,
    `hash` text NOT NULL,
    `client_id` text,
    `scopes` text,
    `expires_at` datetime NOT NULL,
    `revoked` numeric NOT NULL DEFAULT false,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_refresh_tokens_identity` FOREIGN KEY (`identity_id`) REFERENCES `identities`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX `idx_refresh_tokens_expires_at` ON `refresh_tokens`(`expires_at`);
CREATE UNIQUE INDEX `idx_refresh_tokens_hash` ON `refresh_tokens`(`hash`);
CREATE INDEX `idx_refresh_tokens_identity_id` ON `refresh_tokens`(`identity_id`);
-- invitations definition
CREATE TABLE `invitations` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `email` text,
    `token_hash` text,
    `state` text,
    `role` text,
    `created_at` datetime NOT NULL,
    `valid_until` datetime NOT NULL,
    CONSTRAINT `uni_invitations_email` UNIQUE (`email`)
);
CREATE UNIQUE INDEX `idx_invitations_token_hash` ON `invitations`(`token_hash`);
--- PROJECT ---
-- projects definition
CREATE TABLE `projects` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `title` text NOT NULL,
    `created_at` datetime,
    `updated_at` datetime
);
-- Boards table
CREATE TABLE `boards` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `title` text NOT NULL,
    `project_id` integer NOT NULL,
    `created_at` datetime,
    `updated_at` datetime,
    CONSTRAINT `fk_projects_boards` FOREIGN KEY (`project_id`) REFERENCES `projects`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX `idx_boards_project_id` ON `boards`(`project_id`);
-- Columns
CREATE TABLE `columns` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `title` text NOT NULL,
    `order` integer,
    `board_id` integer NOT NULL,
    `project_id` integer NOT NULL,
    `created_at` datetime,
    `updated_at` datetime,
    CONSTRAINT `fk_columns_project` FOREIGN KEY (`project_id`) REFERENCES `projects`(`id`),
    CONSTRAINT `fk_boards_columns` FOREIGN KEY (`board_id`) REFERENCES `boards`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX `idx_columns_project_id` ON `columns`(`project_id`);
CREATE INDEX `idx_columns_board_id` ON `columns`(`board_id`);
-- tasks definition
CREATE TABLE `tasks` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `title` text NOT NULL,
    `order` integer,
    `done` numeric,
    `body` text,
    `created_at` datetime,
    `updated_at` datetime,
    `column_id` integer NOT NULL,
    `project_id` integer NOT NULL,
    `board_id` integer NOT NULL,
    CONSTRAINT `fk_tasks_project` FOREIGN KEY (`project_id`) REFERENCES `projects`(`id`),
    CONSTRAINT `fk_tasks_board` FOREIGN KEY (`board_id`) REFERENCES `boards`(`id`),
    CONSTRAINT `fk_columns_tasks` FOREIGN KEY (`column_id`) REFERENCES `columns`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX `idx_tasks_board_id` ON `tasks`(`board_id`);
CREATE INDEX `idx_tasks_project_id` ON `tasks`(`project_id`);
CREATE INDEX `idx_tasks_column_id` ON `tasks`(`column_id`);
--- OTHER ---
-- activity_logs
CREATE TABLE `activity_logs` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `summary` text NOT NULL,
    `identity_id` integer NOT NULL,
    `entity_id` integer NOT NULL,
    `entity_type` text NOT NULL,
    `action` text NOT NULL,
    `created_at` datetime,
    CONSTRAINT `fk_activity_logs_identity` FOREIGN KEY (`identity_id`) REFERENCES `identities`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX `idx_activity_logs_action` ON `activity_logs`(`action`);
CREATE INDEX `idx_activity_logs_entity_type` ON `activity_logs`(`entity_type`);
CREATE INDEX `idx_activity_logs_entity_id` ON `activity_logs`(`entity_id`);
CREATE INDEX `idx_activity_logs_identity_id` ON `activity_logs`(`identity_id`);
-- jobs definition
CREATE TABLE `jobs` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `type` text,
    `payload` text,
    `status` text,
    `attempts` integer,
    `available_at` datetime,
    `locked_at` datetime,
    `locked_by` text,
    `created_at` datetime
);
