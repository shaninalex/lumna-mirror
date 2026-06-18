package services

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

type WorkspaceManager interface {
	Get(ctx context.Context, id uint) (*models.Workspace, error)
	Create(ctx context.Context, title string) (*models.Workspace, error)
	CreateWithOwner(ctx context.Context, title string, idn *models.Identity) (*models.Workspace, error)
	Update(ctx context.Context, id uint, payload map[string]any) error
	List(ctx context.Context, params map[string]any) ([]*models.Workspace, error)
}

type workspaceService struct {
	repository repositories.WorkspaceRepository
	bus        observer.Observer
}

func NewWorkspaceService(
	repository repositories.WorkspaceRepository,
	bus observer.Observer,
) WorkspaceManager {
	return &workspaceService{
		repository: repository,
		bus:        bus,
	}
}

func (s *workspaceService) Create(ctx context.Context, title string) (*models.Workspace, error) {
	workspace := &models.Workspace{
		Title:     title,
		CreatedAt: time.Now(),
		Active:    true,
		// OwnerEmail: ,
	}

	if err := s.repository.Create(ctx, workspace); err != nil {
		return nil, err
	}

	s.bus.Publish(ctx, models.EventWorkspaceCreated, workspace)

	return workspace, nil
}

func (s *workspaceService) CreateWithOwner(ctx context.Context, title string, idn *models.Identity) (*models.Workspace, error) {
	workspace := &models.Workspace{
		Title:      title,
		CreatedAt:  time.Now(),
		Active:     true,
		OwnerEmail: idn.Email,
	}

	if err := s.repository.Create(ctx, workspace); err != nil {
		return nil, err
	}

	s.bus.Publish(ctx, models.EventWorkspaceCreated, workspace)

	return workspace, nil
}

func (s *workspaceService) Get(ctx context.Context, id uint) (*models.Workspace, error) {
	return s.repository.GetByID(ctx, id)

}

func (s *workspaceService) Update(ctx context.Context, id uint, payload map[string]any) error {
	return s.repository.Update(ctx, id, payload)
}

func (s *workspaceService) List(ctx context.Context, params map[string]any) ([]*models.Workspace, error) {
	return s.repository.List(ctx, params)
}
