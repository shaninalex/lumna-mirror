package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) List(ctx context.Context) ([]models.Project, error) {
	database := db.GetDB(ctx)
	projects := []models.Project{}
	if result := database.Find(&projects); result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}
