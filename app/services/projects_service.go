package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

type ProjectService interface {
	Get(ctx context.Context, id int) (*models.Project, error)
	List(ctx context.Context, workspaceID int) ([]models.Project, error)
	Create(ctx context.Context, data models.ProjectCreateModel) (*models.Project, error)
	Update(ctx context.Context, project *models.Project) error
	Delete(ctx context.Context, id int) error
}

type projectService struct {
	repository     repositories.ProjectRepository
	taskRepository repositories.TaskRepository
	bus            observer.Observer
}

var _ ProjectService = (*projectService)(nil)

func NewProjectService(
	repository repositories.ProjectRepository,
	bus observer.Observer,
) ProjectService {
	s := &projectService{
		repository: repository,
		bus:        bus,
	}
	s.init()
	return s
}

func (s *projectService) init() {
}

func (s *projectService) Get(ctx context.Context, id int) (*models.Project, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *projectService) List(ctx context.Context, workspaceId int) ([]models.Project, error) {
	return s.repository.List(ctx, models.Project{WorkspaceID: workspaceId})
}

func (s *projectService) Create(ctx context.Context, data models.ProjectCreateModel) (*models.Project, error) {
	project := &models.Project{
		Title:       data.Title,
		OwnerID:     data.OwnerID,
		WorkspaceID: data.WorkspaceID,
		Meta:        models.ProjectMeta{},
	}
	if err := s.repository.Create(ctx, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *projectService) Update(ctx context.Context, project *models.Project) error {
	return s.repository.Update(ctx, project)
}

func (s *projectService) Delete(ctx context.Context, id int) error {
	return s.repository.DeleteByID(ctx, id)
}
