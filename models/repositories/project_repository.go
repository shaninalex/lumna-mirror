// Copyright © 2025 Lumna. All rights reserved.

package repositories

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/flowreon/models"
)

// ProjectGetByUserIDAndCode get project by userID and project code
func ProjectGetByUserIDAndCode(ctx context.Context, db *sql.DB, code string) (*models.Project, error) {
	var project models.Project
	query := `SELECT id, user_id, title, code, created_at, updated_at FROM projects WHERE code = ?`
	row := db.QueryRowContext(ctx, query, code)
	if err := row.Scan(&project.ID, &project.UserID, &project.Title, &project.Code, &project.CreatedAt, &project.UpdatedAt); err != nil {
		return nil, err
	}
	return &project, nil
}

// ProjectSave save project
func ProjectSave(ctx context.Context, db *sql.DB, project *models.Project) error {
	query := `INSERT INTO projects (user_id, title, code) VALUES (?, ?, ?)`
	result, err := db.ExecContext(ctx, query, &project.UserID, &project.Title, &project.Code)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	project.SetID(uint(id))
	return err
}

// ProjectList projects list
func ProjectList(ctx context.Context, db *sql.DB) ([]*models.Project, error) {
	query := `SELECT id, user_id, title, code, created_at, updated_at FROM projects`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			panic(err)
		}
	}()
	projects := make([]*models.Project, 0)
	for rows.Next() {
		project := &models.Project{}
		if err = rows.Scan(&project.ID, &project.UserID, &project.Title, &project.Code, &project.CreatedAt, &project.UpdatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}
