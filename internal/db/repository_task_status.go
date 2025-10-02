// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"context"
	"database/sql"
)

// TaskStatusByID get task status by id
func TaskStatusByID(ctx context.Context, db *sql.DB, id uint) (*TaskStatus, error) {
	q := `select id, project_id, title, completed, list_index, config from statuses where id = ?`
	row := db.QueryRowContext(ctx, q, id)
	s := &TaskStatus{}
	if err := row.Scan(&s.ID, &s.ProjectID, &s.Title, &s.Completed, &s.ListIndex, &s.Config); err != nil {
		return nil, err
	}
	return s, nil
}

// TaskStatusListByProject get task status list by project code
func TaskStatusListByProject(ctx context.Context, db *sql.DB, id uint) ([]*TaskStatus, error) {
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
	statuses := []*TaskStatus{}
	for rows.Next() {
		s := &TaskStatus{}
		if err = rows.Scan(&s.ID, &s.ProjectID, &s.Title, &s.Completed, &s.ListIndex, &s.Config); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}
	return statuses, nil
}
