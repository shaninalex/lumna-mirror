package domain

import (
	"context"
	"database/sql"
	"errors"
)

// ProjectGetByID get project by userID and project code
func ProjectGetByID(ctx context.Context, db *sql.DB, id int64) (*Project, error) {
	var project Project
	query := `SELECT id, user_id, title, code, created_at, updated_at FROM projects WHERE id = ?`
	row := db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&project.Id, 0, &project.Title, &project.Code, &project.CreatedAt, &project.UpdatedAt); err != nil {
		return nil, err
	}
	return &project, nil
}

// ProjectSave save project
func ProjectSave(ctx context.Context, db *sql.DB, project *Project) error {
	query := `INSERT INTO projects (user_id, title, code) VALUES (?, ?, ?)`
	result, err := db.ExecContext(ctx, query, &project.UserId, &project.Title, &project.Code)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	project.Id = id
	return err
}

// ProjectList projects list
func ProjectList(ctx context.Context, db *sql.DB) ([]*Project, error) {
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
	projects := make([]*Project, 0)
	for rows.Next() {
		project := &Project{}
		if err = rows.Scan(&project.Id, &project.UserId, &project.Title, &project.Code, &project.CreatedAt, &project.UpdatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func ProjectUpdate(ctx context.Context, db *sql.DB, project *Project) error {
	query := `
		UPDATE 
		    projects
		SET 
			title = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	result, err := db.ExecContext(ctx, query, project.Title, project.Id)
	if err != nil {
		return err
	}
	f, err := result.RowsAffected()
	if f == 0 || err != nil {
		return errors.New("no rows affected")
	}
	return nil
}

func ProjectDelete(ctx context.Context, db *sql.DB, id int64) error {
	query := `
		DELETE FROM 
			projects
		WHERE 
		    id = ?
	`
	res, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
