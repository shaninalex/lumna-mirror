package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
)

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) Get(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	database := persistence.GetDB(ctx)
	project := &models.Project{}
	if result := database.Preload("Boards").Where("id = ?", id.String()).First(&project); result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (s *ProjectService) List(ctx context.Context) ([]models.Project, error) {
	database := persistence.GetDB(ctx)
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
	if result := persistence.GetDB(ctx).Create(&project); result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func (s *ProjectService) Update(ctx context.Context, projectID uuid.UUID, title string) (*models.Project, error) {
	database := persistence.GetDB(ctx)
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
	database := persistence.GetDB(ctx)
	if result := database.Where("id = ?", id).Delete(&models.Project{}); result.Error != nil {
		return result.Error
	}
	return nil
}
