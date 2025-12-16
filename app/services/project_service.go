package services

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type ProjectManager interface {
	List(ctx context.Context) ([]*models.Project, error)
	Create(ctx context.Context, entry *models.Project) error
}

type ProjectService struct {
	projectRepository *repositories.ProjectRespository
}

func NewProjectManager() ProjectManager {
	return &ProjectService{
		projectRepository: repositories.NewProjectRespository(),
	}
}

var _ ProjectManager = (*ProjectService)(nil)

func (s *ProjectService) List(ctx context.Context) ([]*models.Project, error) {
	return s.projectRepository.List(ctx)
}

func (s *ProjectService) Create(ctx context.Context, entry *models.Project) error {
	if entry.Name == "" {
		return fmt.Errorf("project name is required")
	}
	count, err := s.projectRepository.Count(ctx, db.Option{Key: "name", Value: entry.Name})
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("project with name %s already exist", entry.Name)
	}
	return s.projectRepository.Create(ctx, entry)
}
