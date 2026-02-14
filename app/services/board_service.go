package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type BoardService struct {
	db *gorm.DB
}

func NewBoardService(db *gorm.DB) *BoardService {
	return &BoardService{
		db: db,
	}
}

func (s *BoardService) Create(ctx context.Context, projectID uint, title string) (*models.Board, error) {
	board := models.Board{
		ProjectID: projectID,
		Title:     title,
	}
	if result := s.db.WithContext(ctx).Create(&board); result.Error != nil {
		return nil, result.Error
	}
	return &board, nil
}

func (s *BoardService) Delete(ctx context.Context, boardID uint) error {
	if result := s.db.WithContext(ctx).Where("id = ?", boardID).Delete(&models.Board{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *BoardService) Update(ctx context.Context, board *models.Board) error {
	return s.db.WithContext(ctx).Model(&models.Board{}).Save(board).Error
}

func (s *BoardService) Get(ctx context.Context, boardID uint) (*models.Board, error) {
	board := &models.Board{}
	if result := s.db.WithContext(ctx).
		Preload("Columns", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Preload("Columns.Tasks", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Where("id = ?", boardID).
		First(&board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

type KTask struct {
	Id    string `json:"id"`
	Order int    `json:"order"`
}

type KColumn struct {
	Id    string  `json:"id"`
	Order *int    `json:"order"`
	Tasks []KTask `json:"tasks"`
}

type KanbanBoardChangeOrderPayload struct {
	MoveType string    `json:"moveType"`
	ColumnId *string   `json:"columnId"`
	Tasks    []KTask   `json:"tasks"`
	Columns  []KColumn `json:"columns"`
}

func (s *BoardService) Reoder(ctx context.Context, payload *KanbanBoardChangeOrderPayload) error {
	if payload.MoveType == "task" {
		// move in a single column
		if payload.ColumnId != nil {
			for _, task := range payload.Tasks {
				if result := s.db.Model(&models.Task{}).Where("id = ?", task.Id).Update("order", task.Order); result.Error != nil {
					return result.Error
				}
			}
		} else {
			// move task between columns
			for _, column := range payload.Columns {
				for _, task := range column.Tasks {
					if result := s.db.Model(&models.Task{}).
						Where("id = ?", task.Id).
						Updates(map[string]any{
							"order":     task.Order,
							"column_id": column.Id,
						}); result.Error != nil {
						return result.Error
					}
				}
			}
		}
	} else {
		// move column between each other
		for _, column := range payload.Columns {
			if result := s.db.Model(&models.Column{}).Where("id = ?", column.Id).Update("order", column.Order); result.Error != nil {
				return result.Error
			}
		}
	}

	return nil
}
