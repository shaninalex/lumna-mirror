// Copyright © 2025 Lumna. All rights reserved.

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

// TaskStatusListByProject get task status list by project code
func TaskStatusListByProject(ctx context.Context, db *sql.DB, id uint) ([]*models.TaskStatus, error) {
	q := `
	select s.id, s.project_id, s.title, s.completed, s.list_index, s.config 
	from statuses s
	join projects p on p.id = s.project_id                                              
	where p.id = ?
	`
	rows, err := db.QueryContext(ctx, q, id)
	if err != nil {
		return nil, err
	}
	statuses := []*models.TaskStatus{}
	for rows.Next() {
		s := &models.TaskStatus{}
		if err = rows.Scan(&s.ID, &s.ProjectID, &s.Title, &s.Completed, &s.ListIndex, &s.Config); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}
	return statuses, nil
}
