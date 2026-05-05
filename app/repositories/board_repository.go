package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type BoardRepository interface {
	GetByID(ctx context.Context, id uint) (*models.Board, error)
	Create(ctx context.Context, board *models.Board) (*models.Board, error)
	Update(ctx context.Context, board *models.Board) (*models.Board, error)
	Delete(ctx context.Context, id uint) error
}
type GormBoardRepository struct {
	db *gorm.DB
}

func NewGormBoardRepository(db *gorm.DB) BoardRepository {
	return &GormBoardRepository{
		db: db,
	}
}

func (s *GormBoardRepository) GetByID(ctx context.Context, id uint) (*models.Board, error) {
	board := &models.Board{}
	if result := s.db.WithContext(ctx).
		Preload("Columns", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Preload("Columns.Tasks", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Where("id = ?", id).
		First(&board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormBoardRepository) Create(ctx context.Context, board *models.Board) (*models.Board, error) {
	if result := s.db.WithContext(ctx).Create(board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormBoardRepository) Update(ctx context.Context, board *models.Board) (*models.Board, error) {
	if result := s.db.WithContext(ctx).Save(board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *GormBoardRepository) Delete(ctx context.Context, id uint) error {
	if result := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Board{}); result.Error != nil {
		return result.Error
	}
	return nil
}
