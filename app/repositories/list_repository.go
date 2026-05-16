package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ListRepository interface {
	GetByID(ctx context.Context, id uint) (*models.List, error)
	Create(ctx context.Context, board *models.List) (*models.List, error)
	Update(ctx context.Context, board *models.List) (*models.List, error)
	Delete(ctx context.Context, id uint) error
	ListByProjectId(ctx context.Context, projectId uint) ([]*models.List, error) // TODO: make proper filter request. Or not, not sure if it needed
}
type GormListRepository struct {
	db *gorm.DB
}

func NewGormListRepository(db *gorm.DB) ListRepository {
	return &GormListRepository{
		db: db,
	}
}

func (s *GormListRepository) GetByID(ctx context.Context, id uint) (*models.List, error) {
	board := &models.List{}
	if result := s.db.WithContext(ctx).
		Where("id = ?", id).
		First(&board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormListRepository) Create(ctx context.Context, board *models.List) (*models.List, error) {
	if result := s.db.WithContext(ctx).Create(board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormListRepository) Update(ctx context.Context, board *models.List) (*models.List, error) {
	if result := s.db.WithContext(ctx).Save(board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormListRepository) Delete(ctx context.Context, id uint) error {
	if result := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.List{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *GormListRepository) ListByProjectId(ctx context.Context, projectId uint) ([]*models.List, error) {
	boards := make([]*models.List, 0)
	if result := s.db.WithContext(ctx).
		Where("project_id = ?", projectId).
		Find(&boards); result.Error != nil {
		return nil, result.Error
	}
	return boards, nil
}
