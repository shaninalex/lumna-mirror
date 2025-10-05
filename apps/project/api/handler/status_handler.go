// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// ProjectStatusHandler - task handler.
type ProjectStatusHandler struct {
	statusService domain.StatusManager
}

// NewProjectStatusHandler - new task handler.
func NewProjectStatusHandler() *ProjectStatusHandler {
	return &ProjectStatusHandler{
		statusService: domain.NewStatusService(),
	}
}

// Get - Retrieve statuses for a project
func (s *ProjectStatusHandler) Get(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	statuses, err := s.statusService.ProjectStatuses(r.Context(), uint(projectID))
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	web.Success(w, adapter.ToTaskStatusesDto(statuses))
}

// Post - Retrieve statuses for a project
func (s *ProjectStatusHandler) Post(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	payload, err := web.BodyParser[adapter.TaskStatusInput](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	status, err := s.statusService.Create(r.Context(), uint(projectID), payload.Title, payload.Complete)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewTaskStatusDto(status))
}

// Patch - Update statuses for a project
func (s *ProjectStatusHandler) Patch(w http.ResponseWriter, r *http.Request) {
	statusID := web.UrlNumericParam(w, r, "statusId")
	payload, err := web.BodyParser[adapter.TaskStatusInput](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	status, err := s.statusService.Get(ctx, uint(statusID))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	status.Title = payload.Title
	status, err = s.statusService.Patch(ctx, status)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewTaskStatusDto(status))
}

// Delete - Delete status for the project
func (s *ProjectStatusHandler) Delete(w http.ResponseWriter, r *http.Request) {
	statusID := web.UrlNumericParam(w, r, "statusId")
	err := s.statusService.Delete(r.Context(), uint(statusID))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "status deleted")
}
