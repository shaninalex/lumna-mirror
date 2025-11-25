PRAGMA foreign_keys = ON;

------------------------------------------
-- USERS (3)
------------------------------------------
INSERT INTO users (id, email, active)
VALUES (1, 'alice@example.com', 1),
       (2, 'bob@example.com', 1),
       (3, 'charlie@example.com', 1);

------------------------------------------
-- PROJECTS (5)
------------------------------------------
INSERT INTO projects (id, name)
VALUES (1, 'Project Alpha'),
       (2, 'Project Beta'),
       (3, 'Project Gamma'),
       (4, 'Project Delta'),
       (5, 'Project Omega');

------------------------------------------
-- BOARDS (5) – one per project
------------------------------------------
INSERT INTO boards (id, project_id, settings)
VALUES (1, 1, '{}'),
       (2, 2, '{}'),
       (3, 3, '{}'),
       (4, 4, '{}'),
       (5, 5, '{}');

-- Clear old lists if needed
-- DELETE FROM lists;
INSERT INTO lists (board_id, name, "order")
VALUES
-- Board 1 (Project 1)
(1, 'Backlog', 1),
(1, 'In Progress', 2),
(1, 'Done', 3),

-- Board 2 (Project 2)
(2, 'Backlog', 1),
(2, 'In Progress', 2),
(2, 'Done', 3),

-- Board 3 (Project 3)
(3, 'Backlog', 1),
(3, 'In Progress', 2),
(3, 'Done', 3),

-- Board 4 (Project 4)
(4, 'Backlog', 1),
(4, 'In Progress', 2),
(4, 'Done', 3),

-- Board 5 (Project 5)
(5, 'Backlog', 1),
(5, 'In Progress', 2),
(5, 'Done', 3);


------------------------------------------
-- CALENDARS (5) – one per project
------------------------------------------
INSERT INTO calendars (id, project_id, settings)
VALUES (1, 1, '{}'),
       (2, 2, '{}'),
       (3, 3, '{}'),
       (4, 4, '{}'),
       (5, 5, '{}');

------------------------------------------
-- CALENDAR ENTRIES (optional, none added)
------------------------------------------
-- (Skipped – you didn't ask for any)

------------------------------------------
-- TASKS (20)
-- Distributed across the 5 lists (4 each)
------------------------------------------

INSERT INTO tasks (id, board_id, list_id, name, done, created_at, updated_at)
VALUES
-- List 1 (project 1)
(1, 1, 1, 'Task A1', 0, datetime('now'), datetime('now')),
(2, 1, 1, 'Task A2', 1, datetime('now'), datetime('now')),
(3, 1, 1, 'Task A3', 0, datetime('now'), datetime('now')),
(4, 1, 1, 'Task A4', 0, datetime('now'), datetime('now')),

-- List 2 (project 2)
(5, 2, 2, 'Task B1', 0, datetime('now'), datetime('now')),
(6, 2, 2, 'Task B2', 0, datetime('now'), datetime('now')),
(7, 2, 2, 'Task B3', 1, datetime('now'), datetime('now')),
(8, 2, 2, 'Task B4', 0, datetime('now'), datetime('now')),

-- List 3 (project 3)
(9, 3, 3, 'Task C1', 1, datetime('now'), datetime('now')),
(10, 3, 3, 'Task C2', 0, datetime('now'), datetime('now')),
(11, 3, 3, 'Task C3', 0, datetime('now'), datetime('now')),
(12, 3, 3, 'Task C4', 0, datetime('now'), datetime('now')),

-- List 4 (project 4)
(13, 4, 4, 'Task D1', 0, datetime('now'), datetime('now')),
(14, 4, 4, 'Task D2', 0, datetime('now'), datetime('now')),
(15, 4, 4, 'Task D3', 1, datetime('now'), datetime('now')),
(16, 4, 4, 'Task D4', 0, datetime('now'), datetime('now')),

-- List 5 (project 5)
(17, 5, 5, 'Task E1', 0, datetime('now'), datetime('now')),
(18, 5, 5, 'Task E2', 0, datetime('now'), datetime('now')),
(19, 5, 5, 'Task E3', 0, datetime('now'), datetime('now')),
(20, 5, 5, 'Task E4', 1, datetime('now'), datetime('now'));



INSERT INTO comments (id, date, content, entity_type, entity_id, author_id, created_at, updated_at)
VALUES (1, datetime('now'), 'Looks good, please continue.', 'task', 1, 1, datetime('now'), datetime('now')),
       (2, datetime('now'), 'Need to check this later.', 'task', 2, 2, datetime('now'), datetime('now')),
       (3, datetime('now'), 'Blocked by missing data.', 'task', 3, 3, datetime('now'), datetime('now')),
       (4, datetime('now'), 'Can we prioritize this?', 'task', 4, 1, datetime('now'), datetime('now')),

       (5, datetime('now'), 'Great progress so far.', 'task', 5, 2, datetime('now'), datetime('now')),
       (6, datetime('now'), 'Reviewed and approved.', 'task', 6, 3, datetime('now'), datetime('now')),
       (7, datetime('now'), 'Needs a small fix.', 'task', 7, 1, datetime('now'), datetime('now')),
       (8, datetime('now'), 'I will take this one.', 'task', 8, 2, datetime('now'), datetime('now')),

       (9, datetime('now'), 'Waiting for confirmation.', 'task', 9, 3, datetime('now'), datetime('now')),
       (10, datetime('now'), 'Assigned to QA.', 'task', 10, 1, datetime('now'), datetime('now')),
       (11, datetime('now'), 'Added additional details.', 'task', 11, 2, datetime('now'), datetime('now')),
       (12, datetime('now'), 'This might take a bit longer.', 'task', 12, 3, datetime('now'), datetime('now')),

       (13, datetime('now'), 'Please review ASAP.', 'task', 13, 1, datetime('now'), datetime('now')),
       (14, datetime('now'), 'I tested this, looks fine.', 'task', 14, 2, datetime('now'), datetime('now')),
       (15, datetime('now'), 'Could use better naming.', 'task', 15, 3, datetime('now'), datetime('now')),
       (16, datetime('now'), 'Resolved merge conflicts.', 'task', 16, 1, datetime('now'), datetime('now')),

       (17, datetime('now'), 'Documentation updated.', 'task', 17, 2, datetime('now'), datetime('now')),
       (18, datetime('now'), 'I need clarification here.', 'task', 18, 3, datetime('now'), datetime('now')),
       (19, datetime('now'), 'Ready for deployment.', 'task', 19, 1, datetime('now'), datetime('now')),
       (20, datetime('now'), 'Scheduled for next sprint.', 'task', 20, 2, datetime('now'), datetime('now'));
