package pm

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
	"gitlab.com/shaninalex/jajirra/models"
	"gorm.io/gorm"
)

type ProjectManager interface {
	Project(ctx context.Context, orgID uuid.UUID, projectKey string) (*models.Project, error)
	List(ctx context.Context, orgID uuid.UUID) ([]*models.Project, error)
	Issues(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*models.Task, error)
	Statuses(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*models.TaskStatus, error)
}

type ProjectManagement struct{}

func NewProjectManagement() *ProjectManagement {
	return &ProjectManagement{}
}

func (s *ProjectManagement) Project(ctx context.Context, orgID uuid.UUID, projectKey string) (*models.Project, error) {
	var project models.Project

	tx := database.GetDB(ctx).
		WithContext(ctx).
		Preload("Statuses.Tasks").
		Preload("Tasks").
		First(&project, "project_key = ? AND organization_id = ?", projectKey, orgID)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ProjectNotFound
		}
		return nil, tx.Error
	}

	return &project, nil
}

func (s *ProjectManagement) List(ctx context.Context, orgID uuid.UUID) ([]*models.Project, error) {
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

func (s *ProjectManagement) Issues(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*models.Task, error) {
	//var issues []*models.Task
	project, err := s.Project(ctx, orgID, projectKey)
	if err != nil {
		return nil, err
	}
	//tx := database.GetDB(ctx).Where("project_id = ?", project.ID).Find(&issues)
	//if tx.Error != nil {
	//	return nil, tx.Error
	//}
	return project.Tasks, nil
}

func (s *ProjectManagement) Statuses(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*models.TaskStatus, error) {
	project, err := s.Project(ctx, orgID, projectKey)
	if err != nil {
		return nil, err
	}
	return project.Statuses, nil
}
