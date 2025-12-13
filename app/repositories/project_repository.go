package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type ProjectRespository struct {
}

func NewProjectRespository() *ProjectRespository {
	return &ProjectRespository{}
}

var _ Repository[models.Project] = (*ProjectRespository)(nil)

// Count implements Repository.
func (p *ProjectRespository) Count(ctx context.Context, opts ...db.Option) (int, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT count(*) FROM projects %s`,
		where,
	)

	var count int
	row := db.FromContext(ctx).QueryRowContext(ctx, query, args...)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// Create implements Repository.
func (p *ProjectRespository) Create(ctx context.Context, entry *models.Project) error {
	query := `
		INSERT INTO projects (name)
		VALUES (?)
	`

	result, err := db.FromContext(ctx).ExecContext(ctx, query, entry.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	entry.SetId(uint(id))
	return nil
}

// Delete implements Repository.
func (p *ProjectRespository) Delete(ctx context.Context, id uint) error {
	res, err := db.FromContext(ctx).ExecContext(ctx, `DELETE FROM projects WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}

// Get implements Repository.
func (p *ProjectRespository) Get(ctx context.Context, id uint) (*models.Project, error) {
	project := &models.Project{}
	row := db.FromContext(ctx).QueryRow(`
		SELECT id, name FROM projects WHERE id = ?
	`, id)
	err := row.Scan(&project.Id, &project.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorProjectNotFound
		}
		return nil, err
	}

	return project, nil
}

// List implements Repository.
func (p *ProjectRespository) List(ctx context.Context, opts ...db.Option) ([]*models.Project, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT id, name FROM projects %s`,
		where,
	)

	rows, err := db.FromContext(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []*models.Project{}
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(
			&project.Id,
			&project.Name,
		); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}

// Update implements Repository.
func (p *ProjectRespository) Update(ctx context.Context, userID uint, opts ...db.Option) error {
	if len(opts) == 0 {
		return ErrorUserNoFieldsToUpdate
	}

	// // always update updated_at
	// opts = append(opts, db.Option{
	// 	Key:   "updated_at",
	// 	Value: time.Now(),
	// })

	set, args := db.Set(opts)
	if set == "" {
		return ErrorUserNoFieldsToUpdate
	}

	query := fmt.Sprintf(
		`UPDATE projects %s WHERE id = ?`,
		set,
	)

	args = append(args, userID)

	res, err := db.FromContext(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}
