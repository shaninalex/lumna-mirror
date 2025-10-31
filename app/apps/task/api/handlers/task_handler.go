// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"net/http"

	"github.com/shaninalex/lumna/app/apps/task/adapter"
	"github.com/shaninalex/lumna/app/domain"
	"github.com/shaninalex/lumna/app/internal/web"
)

type TaskHandler struct {
	taskManager domain.TaskManager
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{
		taskManager: domain.NewTaskService(),
	}
}

func (s *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	taskID := web.UrlNumericParam(w, r, "id")
	task, err := s.taskManager.TaskDetail(r.Context(), taskID)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.ToTaskDto(task))
}

func (s *TaskHandler) Patch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID := web.UrlNumericParam(w, r, "id")
	payload, err := web.BodyParser[adapter.TaskDetailInput](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	task, err := s.taskManager.TaskDetail(ctx, taskID)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	task.Title = payload.Title
	task.Completed = payload.Completed
	task.Description = &payload.Description
	task.ListIndex = payload.ListIndex
	task.StatusID = payload.StatusID

	err = s.taskManager.TaskUpdate(ctx, task)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, adapter.ToTaskDto(task), "updated")
}

func (s *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	taskID := web.UrlNumericParam(w, r, "id")
	if err := s.taskManager.TaskDelete(r.Context(), taskID); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "Task deleted")
}
