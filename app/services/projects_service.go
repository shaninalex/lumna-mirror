package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

type ProjectService interface {
	Get(ctx context.Context, id uint) (*models.Project, error)
	List(ctx context.Context, workspaceID uint) ([]models.Project, error)
	Create(ctx context.Context, data models.ProjectCreateModel) (*models.Project, error)
	Update(ctx context.Context, project *models.Project) error
	Delete(ctx context.Context, id uint) error
	GenerateKey(ctx context.Context, projectName string) (string, error)
	GetNewCode(ctx context.Context, projectID uint, entity string) (string, error)
}

type projectService struct {
	repository     repositories.ProjectRepository
	taskRepository repositories.TaskRepository
	bus            observer.Observer
}

var _ ProjectService = (*projectService)(nil)

func NewProjectService(
	repository repositories.ProjectRepository,
	taskRepository repositories.TaskRepository,
	bus observer.Observer,
) ProjectService {
	s := &projectService{
		repository:     repository,
		taskRepository: taskRepository,
		bus:            bus,
	}
	s.init()
	return s
}

func (s *projectService) init() {
	s.bus.Subscribe(models.EventTaskCreated, s.handleEventTaskCreated)
}

func (s *projectService) handleEventTaskCreated(ctx context.Context, data any) {
	task, ok := data.(models.Task)
	if !ok {
		return
	}

	project, err := s.Get(ctx, task.ProjectID)
	if err != nil {
		return
	}
	if err = s.Update(ctx, project); err != nil {
		log.Println(err)
	}
}

func (s *projectService) Get(ctx context.Context, id uint) (*models.Project, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *projectService) List(ctx context.Context, workspaceId uint) ([]models.Project, error) {
	return s.repository.List(ctx, models.Project{WorkspaceID: workspaceId})
}

func (s *projectService) Create(ctx context.Context, data models.ProjectCreateModel) (*models.Project, error) {
	key, err := s.GenerateKey(ctx, data.Title)
	if err != nil {
		return nil, err
	}

	project := &models.Project{
		Title:       data.Title,
		OwnerID:     data.OwnerID,
		WorkspaceID: data.WorkspaceID,
		Key:         key,
		Meta:        models.NewProjectMeta(),
	}
	if err := s.repository.Create(ctx, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *projectService) Update(ctx context.Context, project *models.Project) error {
	return s.repository.Update(ctx, project)
}

func (s *projectService) Delete(ctx context.Context, id uint) error {
	return s.repository.DeleteByID(ctx, id)
}

var (
	ErrorProjectUnableGenerateKey = errors.New("unable to generate project key")
	GenerateKeyMaxAttempts        = 25
)

func (s *projectService) GenerateKey(ctx context.Context, projectName string) (string, error) {
	base := utils.ProjectKey(projectName)

	for i := 0; i < GenerateKeyMaxAttempts; i++ {
		key := base

		if i > 0 {
			key = fmt.Sprintf("%s%d", base, i)
		}

		projects, err := s.repository.List(ctx, models.Project{Key: key})
		if err != nil {
			return "", err
		}

		if len(projects) == 0 {
			return key, nil
		}
	}

	return "", ErrorProjectUnableGenerateKey
}

func (s *projectService) GetNewCode(ctx context.Context, projectID uint, entity string) (string, error) {
	project, err := s.repository.GetByID(ctx, projectID)
	if err != nil {
		return "", err
	}

	tasks, err := s.taskRepository.List(ctx, repositories.TaskListQuery{
		ProjectID: &projectID,
		OrderBy:   utils.Pointer("id DESC"),
		Limit:     utils.Pointer(uint(1)),
	})
	if err != nil {
		return "", err
	}

	var n int = 0
	if len(tasks) > 0 {
		n = int(tasks[0].ID)
	}

	return utils.TaskCode(
		project.Title,
		n+1,
	), nil
}
