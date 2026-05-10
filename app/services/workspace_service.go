package services

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type WorkspaceManager interface {
	Get(ctx context.Context, id uint) (*models.Workspace, error)
	Create(ctx context.Context, title string) (*models.Workspace, error)
	Update(ctx context.Context, id uint, payload map[string]any) error
	List(ctx context.Context) ([]*models.Workspace, error)
}

type WorkspaceService struct {
	repository repositories.WorkspaceRepository
}

func NewWorkspaceService(
	repository repositories.WorkspaceRepository,
) WorkspaceManager {
	return &WorkspaceService{
		repository: repository,
	}
}

func (s *WorkspaceService) Create(ctx context.Context, title string) (*models.Workspace, error) {
	workspace := &models.Workspace{
		Title:     title,
		CreatedAt: time.Now(),
	}

	if err := s.repository.Create(ctx, workspace); err != nil {
		return nil, err
	}

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
