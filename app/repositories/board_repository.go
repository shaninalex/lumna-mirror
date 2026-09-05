package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type BoardRepository interface {
	GetByID(ctx context.Context, id int) (*models.Board, error)
	Create(ctx context.Context, board *models.Board) error
	Update(ctx context.Context, board *models.Board) (*models.Board, error)
	Delete(ctx context.Context, id int) error
	ListByProjectId(ctx context.Context, projectId int) ([]*models.Board, error) // TODO: make proper filter request
}

type GormBoardRepository struct {
	db *gorm.DB
}

func NewGormBoardRepository(db *gorm.DB) BoardRepository {
	return &GormBoardRepository{
		db: db,
	}
}

func (s *GormBoardRepository) GetByID(ctx context.Context, id int) (*models.Board, error) {
	board := &models.Board{}
	if result := s.db.WithContext(ctx).
		Where("id = ?", id).
		First(&board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormBoardRepository) Create(ctx context.Context, board *models.Board) error {
	if err := s.db.WithContext(ctx).Create(board).Error; err != nil {
		return err
	}
	return nil
}

func (s *GormBoardRepository) Update(ctx context.Context, board *models.Board) (*models.Board, error) {
	if result := s.db.WithContext(ctx).Save(board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormBoardRepository) Delete(ctx context.Context, id int) error {
	if result := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Board{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *GormBoardRepository) ListByProjectId(ctx context.Context, projectId int) ([]*models.Board, error) {
	boards := make([]*models.Board, 0)
	if result := s.db.WithContext(ctx).
		Where("project_id = ?", projectId).
		Find(&boards); result.Error != nil {
		return nil, result.Error
	}
	return boards, nil
}
