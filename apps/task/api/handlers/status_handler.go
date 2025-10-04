// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/task/adapter"
	"gitlab.com/shaninalex/flowreon/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

type StatusHandler struct {
	taskManager domain.TaskManager
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{
		taskManager: domain.NewTaskService(),
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
	task, err := s.taskManager.TaskDetail(ctx, uint(taskID))
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
