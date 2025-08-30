// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/models"
	"gorm.io/gorm"
)

type ProjectReader interface {
	Project(ctx context.Context, orgID uuid.UUID, projectKey string) (*models.Project, error)
	List(ctx context.Context, orgID uuid.UUID) ([]*models.Project, error)
}

type TaskReader interface {
	TasksList(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*models.TaskStatus, error)
	TaskDetail(ctx context.Context, orgID uuid.UUID, projectKey, taskCode string) (*models.Task, error)
}

type TaskWriter interface {
	PatchTaskStatus(ctx context.Context, orgID uuid.UUID, projectCode string, taskCode string, payload *dto.ChangeTaskStatusDTO) error
	TaskUpdate(ctx context.Context, orgID uuid.UUID, projectKey, taskCode string, data *adapter.UpdateTaskData) error
}

type ProjectManager interface {
	ProjectReader
	TaskReader
	TaskWriter
}

var _ ProjectManager = &ProjectManagement{}

type ProjectManagement struct{}

func NewProjectManagement() *ProjectManagement {
	return &ProjectManagement{}
}

// Project - project.
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

// List - lists all value.
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

// TasksList - tasks list.
func (s *ProjectManagement) TasksList(ctx context.Context, orgID uuid.UUID, projectKey string) ([]*models.TaskStatus, error) {
	project, err := s.Project(ctx, orgID, projectKey)
	if err != nil {
		return nil, err
	}
	return project.Statuses, nil
}

// PatchTaskStatus - patch task status.
func (s *ProjectManagement) PatchTaskStatus(ctx context.Context, orgID uuid.UUID, projectCode string, taskCode string, payload *dto.ChangeTaskStatusDTO) error {
	db := database.GetDB(ctx)
	project, err := s.Project(ctx, orgID, projectCode)
	if err != nil {
		return err
	}
	complete := false
	for _, st := range project.Statuses {
		if st.GetID() == payload.ToStatusID {
			complete = st.Complete
			break
		}
	}

	var task models.Task
	tx := database.GetDB(ctx).Where("code = ? AND project_id = ?", taskCode, project.GetID()).First(&task)
	if tx.Error != nil {
		// TODO: apperror - task not found
		return tx.Error
	}

	task.TaskStatusID = payload.ToStatusID
	task.ListIndex = payload.ToIdx
	task.Completed = complete

	tx = db.Save(&task)
	if tx.Error != nil {
		// TODO: apperror - db error
		return tx.Error
	}
	return nil
}

// TaskDetail - task detail.
func (s *ProjectManagement) TaskDetail(ctx context.Context, orgID uuid.UUID, projectKey, taskCode string) (*models.Task, error) {
	project, err := s.Project(ctx, orgID, projectKey)
	if err != nil {
		return nil, err
	}
	var task models.Task
	tx := database.GetDB(ctx).Where("code = ? AND project_id = ?", taskCode, project.GetID()).First(&task)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &task, nil
}

// TaskUpdate - task update.
func (s *ProjectManagement) TaskUpdate(ctx context.Context, orgID uuid.UUID, projectKey, taskCode string, data *adapter.UpdateTaskData) error {
	db := database.GetDB(ctx)
	project, err := s.Project(ctx, orgID, projectKey)
	if err != nil {
		return err
	}
	var task models.Task
	tx := database.GetDB(ctx).Where("code = ? AND project_id = ?", taskCode, project.GetID()).First(&task)
	if tx.Error != nil {
		return tx.Error
	}
	task.Title = data.Title
	task.Description = data.Description
	tx = db.Save(&task)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
