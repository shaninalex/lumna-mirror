package domain

import (
	"context"
	"database/sql"
	"errors"
)

// TaskStatusByID get task status by id
func TaskStatusByID(ctx context.Context, db *sql.DB, id int64) (*Status, error) {
	q := `select id, project_id, title, list_index, config from statuses where id = ?`
	row := db.QueryRowContext(ctx, q, id)
	s := &Status{}
	if err := row.Scan(&s.Id, &s.ProjectId, &s.Title, &s.ListIndex, &s.Config); err != nil {
		return nil, err
	}
	return s, nil
}

// TaskStatusListByProject get task status list by project code
func TaskStatusListByProject(ctx context.Context, db *sql.DB, id int64) ([]*Status, error) {
	q := `
	select s.id, s.project_id, s.title, s.list_index, s.config 
	from statuses s
	join projects p on p.id = s.project_id                                              
	where p.id = ?
	order by s.list_index
	`
	rows, err := db.QueryContext(ctx, q, id)
	if err != nil {
		return nil, err
	}
	statuses := []*Status{}
	for rows.Next() {
		s := &Status{}
		if err = rows.Scan(&s.Id, &s.ProjectId, &s.Title, &s.ListIndex, &s.Config); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}
	return statuses, nil
}

func TaskStatusCreate(ctx context.Context, db *sql.DB, status *Status) (*Status, error) {
	q := `
	insert into statuses (project_id, title, list_index)
	values (?, ?, ?)
	`
	result, err := db.ExecContext(ctx, q, status.ProjectId, status.Title, status.ListIndex)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("no affected rows")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	status.Id = id
	return status, nil
}

func TaskStatusDelete(ctx context.Context, db *sql.DB, id int64) error {
	q := `delete from statuses where id = ?`
	if _, err := db.ExecContext(ctx, q, id); err != nil {
		return err
	}
	return nil
}

func TaskStatusUpdate(ctx context.Context, db *sql.DB, status *Status) error {
	query := `
		UPDATE statuses
		SET 
			project_id = ?,
			title = ?,
			list_index = ?,
			config = ?
		WHERE id = ?
	`
	_, err := db.ExecContext(ctx, query,
		status.ProjectId,
		status.Title,
		status.ListIndex,
		status.Config,
		status.Id,
	)
	return err
}
