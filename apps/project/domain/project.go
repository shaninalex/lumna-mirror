// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/utils"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/builders"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

// ProjectReader - project reader.
type ProjectReader interface {
	List(ctx context.Context) ([]*models.Project, error)
	GetProject(ctx context.Context, projectCode string) (*models.Project, error)
}

type ProjectWriter interface {
	CreateProject(ctx context.Context, userID uint, projectDto *dto.ProjectDto) (*models.Project, error)
}

type StatusesReader interface {
	ProjectStatuses(ctx context.Context, userID uint, projectCode string) ([]*models.TaskStatus, error)
}

// TaskReader - task reader.
type TaskReader interface {
	TasksList(ctx context.Context, projectCode string) ([]*models.Task, error)
	TaskDetail(ctx context.Context, taskCode string) (*models.Task, error)
}

// TaskWriter - task writer.
type TaskWriter interface {
	PatchTaskStatus(ctx context.Context, taskCode string, payload *dto.ChangeTaskStatusDTO) error
	TaskUpdate(ctx context.Context, taskCode string, data *adapter.UpdateTaskData) error
	TaskCreate(ctx context.Context, userID uint, taskDto *dto.CreateTaskDto) (*models.Task, error)
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

// GetProject - project
func (s *ProjectManagement) GetProject(ctx context.Context, projectCode string) (*models.Project, error) {
	project, err := repositories.ProjectGetByUserIDAndCode(ctx, database.GetDb(ctx), projectCode)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// CreateProject - create new project
func (s *ProjectManagement) CreateProject(ctx context.Context, userID uint, projectDto *dto.ProjectDto) (*models.Project, error) {
	project := builders.NewProjectBuilder().UserID(userID).Title(projectDto.Title).
		Code(utils.GenerateEntityCode("project")).Build()
	if err := repositories.ProjectSave(ctx, database.GetDb(ctx), project); err != nil {
		return nil, err
	}
	return project, nil
}

// List - lists all value.
func (s *ProjectManagement) List(ctx context.Context) ([]*models.Project, error) {
	projects, err := repositories.ProjectList(ctx, database.GetDb(ctx))
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// TasksList - tasks list.
func (s *ProjectManagement) TasksList(ctx context.Context, projectCode string) ([]*models.Task, error) {
	tasks, err := repositories.TaskList(ctx, database.GetDb(ctx), projectCode)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// PatchTaskStatus - patch task status.
func (s *ProjectManagement) PatchTaskStatus(ctx context.Context, taskCode string, payload *dto.ChangeTaskStatusDTO) error {
	db := database.GetDb(ctx)
	task, err := repositories.TaskGet(ctx, db, taskCode)
	if err != nil {
		return err
	}
	task.StatusID = payload.ToStatusID
	task.ListIndex = payload.ToIdx

	status, err := repositories.TaskStatusByID(ctx, db, task.StatusID)
	if err != nil {
		return err
	}
	task.Completed = status.Completed
	err = repositories.UpdateTask(ctx, db, taskCode, task)
	if err != nil {
		return err
	}
	return nil
}

// TaskDetail - task detail.
func (s *ProjectManagement) TaskDetail(ctx context.Context, taskCode string) (*models.Task, error) {
	task, err := repositories.TaskGet(ctx, database.GetDb(ctx), taskCode)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// TaskUpdate - task update.
func (s *ProjectManagement) TaskUpdate(ctx context.Context, taskCode string, data *adapter.UpdateTaskData) error {
	// This is very strange method. Why do not use single method instead of TaskUpdate and PatchTaskStatus ???
	db := database.GetDb(ctx)
	task, err := repositories.TaskGet(ctx, db, taskCode)
	if err != nil {
		return err
	}
	task.Title = data.Title
	task.Description = &data.Description
	err = repositories.UpdateTask(ctx, db, taskCode, task)
	if err != nil {
		return err
	}
	return nil
}

// TaskCreate - create new task.
func (s *ProjectManagement) TaskCreate(ctx context.Context, userID uint, taskDto *dto.CreateTaskDto) (*models.Task, error) {
	db := database.GetDb(ctx)
	project, err := s.GetProject(ctx, taskDto.ProjectCode)
	if err != nil {
		return nil, err
	}
	task := &models.Task{
		UserID:    userID,
		ProjectID: project.GetID(),
		Code:      utils.GenerateEntityCode("task"),
		StatusID:  taskDto.StatusID,
		Title:     taskDto.Title,
		ListIndex: 0,
	}
	err = repositories.TaskSave(ctx, db, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
