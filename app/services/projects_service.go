package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) Get(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	database := db.GetDB(ctx)
	project := &models.Project{}
	if result := database.Preload("Boards").Where("id = ?", id.String()).First(&project); result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (s *ProjectService) GetBoard(ctx context.Context, boardID uuid.UUID) (*models.Board, error) {
	database := db.GetDB(ctx)
	board := &models.Board{}
	if result := database.
		Preload("Lists", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Preload("Lists.Tasks", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("\"order\" ASC")
		}).
		Where("id = ?", boardID.String()).
		First(&board); result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func (s *ProjectService) List(ctx context.Context) ([]models.Project, error) {
	database := db.GetDB(ctx)
	projects := []models.Project{}
	if result := database.Preload("Boards").Find(&projects); result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (s *ProjectService) Create(ctx context.Context, title string, userID uuid.UUID) (*models.Project, error) {
	project := models.Project{
		Title: title,
	}
	database := db.GetDB(ctx)
	if result := database.Create(&project); result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func (s *ProjectService) Update(ctx context.Context, projectID uuid.UUID, title string) (*models.Project, error) {
	database := db.GetDB(ctx)
	project := models.Project{}
	if result := database.Where("id = ?", projectID).First(&project); result.Error != nil {
		return nil, result.Error
	}

	project.Title = title

	if result := database.Save(&project); result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func (s *ProjectService) Delete(ctx context.Context, id uuid.UUID) error {
	database := db.GetDB(ctx)
	if result := database.Where("id = ?", id).Delete(&models.Project{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ProjectService) GetTask(ctx context.Context, taskID uuid.UUID) (*models.Task, error) {
	database := db.GetDB(ctx)
	task := &models.Task{}
	if result := database.Where("id = ?", taskID).First(&task); result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (s *ProjectService) ReorderTask(ctx context.Context, taskID uuid.UUID, boardListID uuid.UUID, order uint) error {
	database := db.GetDB(ctx)
	return database.Model(&models.Task{}).
		Where("id = ?", taskID).
		Updates(map[string]any{
			"board_list_id": boardListID,
			"order":         order,
		}).Error
}

func (s *ProjectService) ReorderList(ctx context.Context, listID uuid.UUID, order uint) error {
	database := db.GetDB(ctx)
	return database.Model(&models.BoardList{}).
		Where("id = ?", listID).
		Update("order", order).Error
}

func (s *ProjectService) UpdateProject(ctx context.Context, projectID uuid.UUID, title string) error {
	database := db.GetDB(ctx)
	return database.Model(&models.Project{}).
		Where("id = ?", projectID.String()).
		Update("title", title).Error
}
