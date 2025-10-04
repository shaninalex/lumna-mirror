// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"context"
	"database/sql"
)

// TaskList task list
func TaskList(ctx context.Context, db *sql.DB, projectID uint) ([]*Task, error) {
	query := `select t.id, t.user_id, t.project_id, t.status_id, t.title, t.code, t.completed, t.description, t.list_index, t.created_at, t.updated_at
		from tasks t
		join projects p on p.id = t.project_id
		where p.id = ?
	`
	rows, err := db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			panic(err)
		}
	}()
	tasks := make([]*Task, 0)
	for rows.Next() {
		t := &Task{}
		if err = rows.Scan(&t.ID, &t.UserID, &t.ProjectID, &t.StatusID, &t.Title, &t.Code, &t.Completed, &t.Description, &t.ListIndex, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// TaskUpdate update task
func TaskUpdate(ctx context.Context, db *sql.DB, id uint, task *Task) error {
	query := `
		UPDATE tasks
		SET 
			title = ?,
			status_id = ?,
			user_id = ?,
			description = ?,
			completed = ?,
			list_index = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	_, err := db.ExecContext(ctx, query,
		task.Title,
		task.StatusID,
		task.UserID,
		task.Description,
		task.Completed,
		task.ListIndex,
		id,
	)
	return err
}

// TaskGet get task
func TaskGet(ctx context.Context, db *sql.DB, id uint) (*Task, error) {
	query := `select t.id, t.user_id, t.project_id, t.status_id, t.title, t.code, t.completed, t.description, t.list_index, t.created_at, t.updated_at
		from tasks t
		where t.id = ?
	`
	row := db.QueryRowContext(ctx, query, id)
	t := &Task{}
	if err := row.Scan(&t.ID, &t.UserID, &t.ProjectID, &t.StatusID, &t.Title, &t.Code, &t.Completed, &t.Description, &t.ListIndex, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}
	return t, nil
}

// TaskSave save task
func TaskSave(ctx context.Context, db *sql.DB, t *Task) error {
	query := `
		insert into tasks (user_id, project_id, status_id, title, code, completed, description, list_index)
		values (?, ?, ?, ?, ?, ?, ?, ?);
	`
	result, err := db.ExecContext(ctx, query, &t.ID, &t.ProjectID, &t.StatusID, &t.Title, &t.Code, &t.Completed, &t.Description, &t.ListIndex)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = uint(id)
	return nil
}

// TaskDelete - delete task by id
func TaskDelete(ctx context.Context, db *sql.DB, id uint) error {
	query := `delete from tasks where id = ?;`
	_, err := db.ExecContext(ctx, query, id)
	return err
}
