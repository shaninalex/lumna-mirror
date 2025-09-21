// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/utils"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/builders"
	"gorm.io/gorm"
)

// ProjectReader - project reader.
type ProjectReader interface {
	List(ctx context.Context, userID uuid.UUID) ([]*models.Project, error)
	Project(ctx context.Context, userID uuid.UUID, projectCode string) (*models.Project, error)
}

type ProjectWriter interface {
	CreateProject(ctx context.Context, userID uuid.UUID, projectDto *dto.ProjectDto) (*models.Project, error)
}

type StatusesReader interface {
	ProjectStatuses(ctx context.Context, userID uuid.UUID, projectCode string) ([]*models.TaskStatus, error)
}

// TaskReader - task reader.
type TaskReader interface {
	TasksList(ctx context.Context, userID uuid.UUID, projectCode string) ([]*models.Task, error)
	TaskDetail(ctx context.Context, userID uuid.UUID, projectCode, taskCode string) (*models.Task, error)
}

// TaskWriter - task writer.
type TaskWriter interface {
	PatchTaskStatus(ctx context.Context, userID uuid.UUID, projectCode string, taskCode string, payload *dto.ChangeTaskStatusDTO) error
	TaskUpdate(ctx context.Context, userID uuid.UUID, projectCode, taskCode string, data *adapter.UpdateTaskData) error
	TaskCreate(ctx context.Context, userID uuid.UUID, taskDto *dto.CreateTaskDto) (*models.Task, error)
}

// ProjectManager - project manager.
type ProjectManager interface {
	ProjectReader
	ProjectWriter
	TaskReader
	TaskWriter
}

var _ ProjectManager = &ProjectManagement{}

// ProjectManagement - project management.
type ProjectManagement struct{}

// NewProjectManagement - new project management.
func NewProjectManagement() *ProjectManagement {
	return &ProjectManagement{}
}

// Project - project. TODO: should be GetProject ( do not forget the verb! )
func (s *ProjectManagement) Project(ctx context.Context, orgID uuid.UUID, projectCode string) (*models.Project, error) {
	var project models.Project
	tx := database.GetDB(ctx).
		WithContext(ctx).
		First(&project, "project_key = ? AND organization_id = ?", projectCode, orgID)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ProjectNotFound
		}
		return nil, tx.Error
	}

	return &project, nil
}

// CreateProject - create new project
func (s *ProjectManagement) CreateProject(ctx context.Context, userID uuid.UUID, projectDto *dto.ProjectDto) (*models.Project, error) {
	project := builders.NewProjectBuilder().
		ID(uuid.New()).
		UserID(userID).
		Title(projectDto.Title).
		Code(utils.GenerateEntityCode("project")).
		Build()
	tx := database.GetDB(ctx).Create(project)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return project, nil
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
func (s *ProjectManagement) TasksList(ctx context.Context, orgID uuid.UUID, projectCode string) ([]*models.Task, error) {
	project, err := s.Project(ctx, orgID, projectCode)
	if err != nil {
		return nil, err
	}
	var tasks []*models.Task
	tx := database.GetDB(ctx).Where("project_id = ?", project.GetID()).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tasks, nil
}

// PatchTaskStatus - patch task status.
func (s *ProjectManagement) PatchTaskStatus(ctx context.Context, orgID uuid.UUID, projectCode, taskCode string, payload *dto.ChangeTaskStatusDTO) error {
	db := database.GetDB(ctx)
	project, err := s.Project(ctx, orgID, projectCode)
	if err != nil {
		return err
	}
	complete := false
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
		return errors.Join(apperrors.TaskUnableToPatch, tx.Error)
	}
	return nil
}

// TaskDetail - task detail.
func (s *ProjectManagement) TaskDetail(ctx context.Context, orgID uuid.UUID, projectCode, taskCode string) (*models.Task, error) {
	project, err := s.Project(ctx, orgID, projectCode)
	if err != nil {
		return nil, err
	}
	var task models.Task
	tx := database.GetDB(ctx).Where("code = ? AND project_id = ?", taskCode, project.GetID()).First(&task)
	if tx.Error != nil {
		return nil, errors.Join(apperrors.TaskNotFound, tx.Error)
	}
	return &task, nil
}

// TaskUpdate - task update.
func (s *ProjectManagement) TaskUpdate(ctx context.Context, orgID uuid.UUID, projectCode, taskCode string, data *adapter.UpdateTaskData) error {
	db := database.GetDB(ctx)
	project, err := s.Project(ctx, orgID, projectCode)
	if err != nil {
		return err
	}
	var task models.Task
	tx := database.GetDB(ctx).Where("code = ? AND project_id = ?", taskCode, project.GetID()).First(&task)
	if tx.Error != nil {
		return errors.Join(apperrors.TaskNotFound, tx.Error)
	}
	task.Title = data.Title
	task.Description = data.Description
	tx = db.Save(&task)
	if tx.Error != nil {
		return errors.Join(apperrors.TaskUnableToPatch, tx.Error)
	}
	return nil
}

// TaskCreate - create new task.
func (s *ProjectManagement) TaskCreate(ctx context.Context, userID uuid.UUID, taskDto *dto.CreateTaskDto) (*models.Task, error) {
	db := database.GetDB(ctx)
	project, err := s.Project(ctx, userID, taskDto.ProjectCode)
	if err != nil {
		return nil, err
	}
	task := models.Task{
		UserID:       userID,
		ProjectID:    project.GetID(),
		Code:         utils.GenerateEntityCode("task"),
		TaskStatusID: taskDto.StatusID,
		Title:        taskDto.Title,
		ListIndex:    0,
	}
	tx := db.Create(&task)
	if tx.Error != nil {
		return nil, errors.Join(apperrors.TaskUnableToPatch, tx.Error)
	}
	return &task, nil
}
