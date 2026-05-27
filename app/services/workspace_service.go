package services

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/observer"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type WorkspaceManager interface {
	Get(ctx context.Context, id uint) (*models.Workspace, error)
	Create(ctx context.Context, title string) (*models.Workspace, error)
	CreateWithOwner(ctx context.Context, title string, idn *models.Identity) (*models.Workspace, error)
	Update(ctx context.Context, id uint, payload map[string]any) error
	List(ctx context.Context) ([]*models.Workspace, error)
}

type WorkspaceService struct {
	repository repositories.WorkspaceRepository
	bus        observer.Observer
}

func NewWorkspaceService(
	repository repositories.WorkspaceRepository,
	bus observer.Observer,
) WorkspaceManager {
	return &WorkspaceService{
		repository: repository,
		bus:        bus,
	}
}

func (s *WorkspaceService) Create(ctx context.Context, title string) (*models.Workspace, error) {
	workspace := &models.Workspace{
		Title:     title,
		Slug:      utils.Slugify(title),
		CreatedAt: time.Now(),
		// OwnerEmail: ,
	}

	if err := s.repository.Create(ctx, workspace); err != nil {
		return nil, err
	}

	s.bus.Publish(ctx, models.EventWorkspaceCreated, workspace)

	return workspace, nil
}

func (s *WorkspaceService) CreateWithOwner(ctx context.Context, title string, idn *models.Identity) (*models.Workspace, error) {
	workspace := &models.Workspace{
		Title:      title,
		Slug:       utils.Slugify(title),
		CreatedAt:  time.Now(),
		OwnerEmail: idn.Email,
	}

	if err := s.repository.Create(ctx, workspace); err != nil {
		return nil, err
	}

	s.bus.Publish(ctx, models.EventWorkspaceCreated, workspace)

	return workspace, nil
}

func (s *WorkspaceService) Get(ctx context.Context, id uint) (*models.Workspace, error) {
	return s.repository.GetByID(ctx, id)

}

func (s *WorkspaceService) Update(ctx context.Context, id uint, payload map[string]any) error {
	return s.repository.Update(ctx, id, payload)
}

func (s *WorkspaceService) List(ctx context.Context) ([]*models.Workspace, error) {
	return s.repository.List(ctx, map[string]any{})
}
