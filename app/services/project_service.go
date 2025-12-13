package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
)

type ProjectManager interface {
	List(ctx context.Context) ([]models.Project, error)
}

type ProjectService struct {
}

var _ ProjectManager = (*ProjectService)(nil)

// List implements ProjectManager.
func (p *ProjectService) List(ctx context.Context) ([]models.Project, error) {
	panic("unimplemented")
}
