// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"net/http"

	"github.com/shaninalex/lumna/app/apps/task/adapter"
	"github.com/shaninalex/lumna/app/domain"
	"github.com/shaninalex/lumna/app/internal/web"
)

type StatusHandler struct {
	taskManager   domain.TaskManager
	statusManager domain.StatusManager
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{
		taskManager:   domain.NewTaskService(),
		statusManager: domain.NewStatusService(),
	}
}

func (s *StatusHandler) Patch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID := web.UrlNumericParam(w, r, "id")
	payload, err := web.BodyParser[adapter.ChangeTaskStatusInput](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	task, err := s.taskManager.TaskDetail(ctx, taskID)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	task.StatusID = payload.ToStatusID
	task.ListIndex = payload.ToIdx
	err = s.taskManager.TaskUpdate(ctx, task)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.ToTaskDto(task), "Status updated")
}
