package pm

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
)

type Projects struct{}

func NewProjects() *Projects {
	return &Projects{}
}

func (s *Projects) List(ctx context.Context, orgID uuid.UUID) ([]*database.Project, error) {
	db := database.GetDB(ctx)
	var projects []*database.Project
	result := db.Find(&projects).Where("organization_id = ?", orgID)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, apperrors.ProjectNotFound
		}
		return nil, result.Error
	}
	return projects, nil
}

func (s *Projects) ProjectTasks(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*database.Issue, error) {
	db := database.GetDB(ctx)
	var issues []*database.Issue
	var project *database.Project
	db.First(&project, "project_key = ? AND organization_id = ?", projectKey, orgID)
	err := db.Model(&project).Association("Issues").Find(&issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
}
