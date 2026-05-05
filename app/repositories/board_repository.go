package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type BoardRepository interface {
	GetByID(ctx context.Context, id uint) (*models.Board, error)
}
type GormBoardRepository struct {
	db *gorm.DB
}

func NewGormBoardRepository(db *gorm.DB) BoardRepository {
	return &GormBoardRepository{
		db: db,
	}
}

func (s GormBoardRepository) GetByID(ctx context.Context, id uint) (*models.Board, error) {
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
