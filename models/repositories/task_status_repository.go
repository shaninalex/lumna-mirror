// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package repositories

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/flowreon/models"
)

// TaskStatusByID get task status by id
func TaskStatusByID(ctx context.Context, db *sql.DB, id uint) (*models.TaskStatus, error) {
	q := `select id, project_id, title, completed, list_index, config from statuses where id = ?`
	row := db.QueryRowContext(ctx, q, id)
	s := &models.TaskStatus{}
	if err := row.Scan(&s.ID, &s.ProjectID, &s.Title, &s.Completed, &s.ListIndex, &s.Config); err != nil {
		return nil, err
	}
	return s, nil
}
