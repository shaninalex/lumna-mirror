// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
	"gitlab.com/shaninalex/flowreon/apps/project/domain/services"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// ProjectTaskHandler - task handler.
type ProjectTaskHandler struct {
	taskManager services.TaskManager
}

// NewProjectTaskHandler - new task handler.
func NewProjectTaskHandler() *ProjectTaskHandler {
	return &ProjectTaskHandler{
		taskManager: services.NewTaskService(),
	}
}

// List - retrieve tasks for a project
func (s *ProjectTaskHandler) List(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	tasks, err := s.taskManager.TasksList(r.Context(), uint(projectID))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.ToTaskListDto(tasks))
}

// Create - create a new task in a project
func (s *ProjectTaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	input, err := web.BodyParser[adapter.ProjectInput](r)
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	ctx := r.Context()
	userID := web.GetUserID(r)

	task := &models.Task{
		Title:     input.Title,
		ProjectID: uint(projectID),
		UserID:    userID,
	}
	task, err = s.taskManager.TaskCreate(ctx, task)
	if err != nil {
		web.Error(w, http.StatusInternalServerError, err)
		return
	}
	web.Success(w, adapter.ToTaskDto(task))
}
