package repositories

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type ProjectRespository struct {
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
	panic("unimplemented")
}

// Delete implements Repository.
func (p *ProjectRespository) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// Get implements Repository.
func (p *ProjectRespository) Get(ctx context.Context, id uint) (*models.Project, error) {
	panic("unimplemented")
}

// List implements Repository.
func (p *ProjectRespository) List(ctx context.Context, opts ...db.Option) ([]*models.Project, error) {
	panic("unimplemented")
}

// Update implements Repository.
func (p *ProjectRespository) Update(ctx context.Context, entry *models.Project, opts ...db.Option) error {
	panic("unimplemented")
}
