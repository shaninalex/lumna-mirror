package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/logger"
)

type ProjectService interface {
	Get(ctx context.Context, id uint) (*models.Project, error)
	List(ctx context.Context) ([]models.Project, error)
	Create(ctx context.Context, data models.ProjectCreateModel) (*models.Project, error)
	Update(ctx context.Context, projectID uint, title string) (*models.Project, error)
	Delete(ctx context.Context, id uint) error
}

type projectService struct {
	repository repositories.ProjectRepository
	logger     logger.Logger
}

var _ ProjectService = (*projectService)(nil)

func NewProjectService(
	repository repositories.ProjectRepository,
	logger logger.Logger,
) ProjectService {
	s := &projectService{
		repository: repository,
		logger:     logger,
	}
	return s
}

func (s *projectService) Get(ctx context.Context, id uint) (*models.Project, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *projectService) List(ctx context.Context) ([]models.Project, error) {
	return s.repository.List(ctx)
}

func (s *projectService) Create(ctx context.Context, data models.ProjectCreateModel) (*models.Project, error) {
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

func (s *projectService) Update(ctx context.Context, projectID uint, title string) (*models.Project, error) {
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

func (s *projectService) Delete(ctx context.Context, id uint) error {
	return s.repository.DeleteByID(ctx, id)
}
