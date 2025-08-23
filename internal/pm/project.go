package pm

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
	"gitlab.com/shaninalex/jajirra/models"
)

type Projects struct{}

func NewProjects() *Projects {
	return &Projects{}
}

func (s *Projects) List(ctx context.Context, orgID uuid.UUID) ([]*models.Project, error) {
	db := database.GetDB(ctx)
	var projects []*models.Project
	result := db.Find(&projects).Where("organization_id = ?", orgID)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, apperrors.ProjectNotFound
		}
		return nil, result.Error
	}
	return projects, nil
}

func (s *Projects) ProjectTasks(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*models.Issue, error) {
	db := database.GetDB(ctx)
	var issues []*models.Issue
	var project *models.Project
	tx := db.First(&project, "project_key = ? AND organization_id = ?", projectKey, orgID)
	if tx.Error != nil {
		if tx.Error.Error() == "record not found" {
			return nil, apperrors.ProjectNotFound
		}
		return nil, tx.Error
	}

	tx = db.Where("project_id = ?", project.ID).Find(&issues)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return issues, nil
}
