package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type WorkspaceRepository interface {
	List(ctx context.Context, query map[string]any) ([]*models.WorkspaceModel, error)
	GetByID(ctx context.Context, taskID uint) (*models.WorkspaceModel, error)
	Create(ctx context.Context, task *models.WorkspaceModel) error
	Update(ctx context.Context, taskID uint, updates map[string]any) error
	// Before implement this method - we should understand what exactly will be deleted
	// ScheduleDelete(ctx context.Context, taskID uint) error
}

type GormWorkspaceRepository struct {
	db *gorm.DB
}

func NewGormWorkspaceRepository(db *gorm.DB) WorkspaceRepository {
	return &GormWorkspaceRepository{db: db}
}

func (s *GormWorkspaceRepository) List(ctx context.Context, query map[string]any) ([]*models.WorkspaceModel, error) {
	var workspaces []*models.WorkspaceModel
	if err := s.db.Where(query).Find(&workspaces).Error; err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (s *GormWorkspaceRepository) GetByID(ctx context.Context, workspaceId uint) (*models.WorkspaceModel, error) {
	var workspace models.WorkspaceModel
	if err := s.db.WithContext(ctx).
		Where("id = ?", workspaceId).
		First(&workspace).
		Error; err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (s *GormWorkspaceRepository) Create(ctx context.Context, workspace *models.WorkspaceModel) error {
	return s.db.WithContext(ctx).
		Create(workspace).
		Error
}

func (s *GormWorkspaceRepository) Update(ctx context.Context, workspaceId uint, updates map[string]any) error {
	result := s.db.WithContext(ctx).Model(&models.WorkspaceModel{}).
		Where("id = ?", workspaceId).
		Updates(updates)
	return result.Error
}
