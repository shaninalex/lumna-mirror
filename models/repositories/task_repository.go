// Copyright © 2025 Lumna. All rights reserved.

package repositories

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/flowreon/models"
)

// TaskList task list
func TaskList(ctx context.Context, db *sql.DB, code string) ([]*models.Task, error) {
	query := `select t.id, t.user_id, t.project_id, t.status_id, t.title, t.code, t.completed, t.description, t.list_index, t.created_at, t.updated_at
		from tasks t
		join projects p on p.id = t.project_id
		where p.code = ?
	`
	rows, err := db.QueryContext(ctx, query, code)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			panic(err)
		}
	}()
	tasks := make([]*models.Task, 0)
	for rows.Next() {
		t := &models.Task{}
		if err = rows.Scan(&t.ID, &t.UserID, &t.ProjectID, &t.StatusID, &t.Title, &t.Code, &t.Completed, &t.Description, &t.ListIndex, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// UpdateTask update task
func UpdateTask(ctx context.Context, db *sql.DB, code string, task *models.Task) error {
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
		WHERE code = ?
	`
	_, err := db.ExecContext(ctx, query,
		task.Title,
		task.StatusID,
		task.UserID,
		task.Description,
		task.Completed,
		task.ListIndex,
		code,
	)
	return err
}

// TaskGet get task
func TaskGet(ctx context.Context, db *sql.DB, code string) (*models.Task, error) {
	query := `select t.id, t.user_id, t.project_id, t.status_id, t.title, t.code, t.completed, t.description, t.list_index, t.created_at, t.updated_at
		from tasks t
		where t.code = ?
	`
	row := db.QueryRowContext(ctx, query, code)
	t := &models.Task{}
	if err := row.Scan(&t.ID, &t.UserID, &t.ProjectID, &t.StatusID, &t.Title, &t.Code, &t.Completed, &t.Description, &t.ListIndex, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}
	return t, nil
}

// TaskSave save task
func TaskSave(ctx context.Context, db *sql.DB, t *models.Task) error {
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
