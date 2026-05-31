package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/logger"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type ProjectService struct {
	repository repositories.ProjectRepository
	logger     logger.Logger
}

func NewProjectService(
	repository repositories.ProjectRepository,
	logger logger.Logger,
) *ProjectService {
	s := &ProjectService{
		repository: repository,
		logger:     logger,
	}
	return s
}

func (s *ProjectService) Get(ctx context.Context, id uint) (*models.Project, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *ProjectService) List(ctx context.Context) ([]models.Project, error) {
	return s.repository.List(ctx)
}

func (s *ProjectService) Create(ctx context.Context, data models.ProjectCreateModel) (*models.Project, error) {
	project := &models.Project{
		Title:       data.Title,
		OwnerID:     data.OwnerID,
		WorkspaceID: data.WorkspaceID,
	}
	if err := s.repository.Create(ctx, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *ProjectService) Update(ctx context.Context, projectID uint, title string) (*models.Project, error) {
	project, err := s.repository.GetByID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	project.Title = title

	if err := s.repository.Update(ctx, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *ProjectService) Delete(ctx context.Context, id uint) error {
	return s.repository.DeleteByID(ctx, id)
}
