package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type BoardService struct {
}

func NewBoardService() *BoardService {
	return &BoardService{}
}

func (s *BoardService) Create(ctx context.Context, projectID uuid.UUID, title string) (*models.Board, error) {
	board := models.Board{
		ProjectID: projectID,
		Title:     title,
	}
	if result := persistence.GetDB(ctx).Create(&board); result.Error != nil {
		return nil, result.Error
	}
	return &board, nil
}

func (s *BoardService) BoardDelete(ctx context.Context, boardID uuid.UUID) error {
	if result := persistence.GetDB(ctx).Where("id = ?", boardID).Delete(&models.Board{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *BoardService) BoardUpdate(ctx context.Context, projectID, boardID uuid.UUID, title string) error {
	database := persistence.GetDB(ctx)
	return database.Model(&models.Board{}).
		Where("id = ? AND project_id = ?", boardID.String(), projectID.String()).
		Update("title", title).Error
}

func (s *BoardService) BoardGet(ctx context.Context, boardID uuid.UUID) (*models.Board, error) {
	database := persistence.GetDB(ctx)
	board := &models.Board{}
	if result := database.
		Preload("Columns", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Preload("Columns.Tasks", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Where("id = ?", boardID.String()).
		First(&board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}
