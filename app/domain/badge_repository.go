package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

// BadgeGet fetches a badge by Id.
func BadgeGet(ctx context.Context, db *sql.DB, id int64) (*Badge, error) {
	row := db.QueryRowContext(ctx, `
		SELECT id, project_id, title, config, created_at
		FROM badge
		WHERE id = ?`, id)

	var b Badge
	if err := row.Scan(&b.Id, &b.ProjectID, &b.Title, &b.Config, &b.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // not found
		}
		return nil, err
	}
	return &b, nil
}

// BadgeCreate inserts a new badge and updates its Id.
func BadgeCreate(ctx context.Context, db *sql.DB, badge *Badge) error {
	b, err := json.Marshal(badge.Config)
	if err != nil {
		return err
	}
	res, err := db.ExecContext(ctx, `
		INSERT INTO badge (project_id, title, config)
		VALUES (?, ?, ?)`,
		badge.ProjectID, badge.Title, string(b))
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	badge.Id = id
	return nil
}

// BadgeUpdate updates an existing badge.
func BadgeUpdate(ctx context.Context, db *sql.DB, badge *Badge) error {
	_, err := db.ExecContext(ctx, `
		UPDATE badge
		SET project_id = ?, title = ?, config = ?
		WHERE id = ?`,
		badge.ProjectID, badge.Title, badge.Config, badge.Id)
	return err
}

// BadgeDelete deletes a badge by Id.
func BadgeDelete(ctx context.Context, db *sql.DB, projectId, id int64) error {
	_, err := db.ExecContext(ctx, `DELETE FROM badge WHERE project_id = ? AND id = ?`, projectId, id)
	return err
}

// BadgeProjectList returns all badges belonging to a project.
func BadgeProjectList(ctx context.Context, db *sql.DB, projectID int64) ([]*Badge, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id, project_id, title, config, created_at
		FROM badge
		WHERE project_id = ?
		ORDER BY created_at ASC`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Badge
	for rows.Next() {
		var b Badge
		var badgeConfig BadgeConfig
		if err := rows.Scan(&b.Id, &b.ProjectID, &b.Title, &badgeConfig, &b.CreatedAt); err != nil {
			return nil, err
		}
		b.Config = badgeConfig
		result = append(result, &b)
	}
	return result, rows.Err()
}

// BadgeTaskList returns all badges assigned to a given task.
func BadgeTaskList(ctx context.Context, db *sql.DB, taskID int64) ([]*Badge, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT b.id, b.project_id, b.title, b.config, b.created_at
		FROM badge b
		JOIN badges_tasks bt ON b.id = bt.badge_id
		WHERE bt.task_id = ?
		ORDER BY b.created_at ASC`, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Badge
	for rows.Next() {
		var b Badge
		if err := rows.Scan(&b.Id, &b.ProjectID, &b.Title, &b.Config, &b.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &b)
	}
	return result, rows.Err()
}
