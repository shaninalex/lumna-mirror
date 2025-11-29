package handler

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/_old_app/apps/project/adapter"
	"gitlab.com/shaninalex/lumna/_old_app/domain"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/web"
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
	statuses, err := s.statusService.ProjectStatuses(r.Context(), projectID)
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
	status, err := s.statusService.Create(r.Context(), projectID, payload.Title)
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
	status, err := s.statusService.Get(ctx, statusID)
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
	err := s.statusService.Delete(r.Context(), statusID)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "status deleted")
}

func (s *ProjectStatusHandler) PatchSort(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	payload, err := web.BodyParser[map[int64]int64](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	err = s.statusService.SortProjectStatus(r.Context(), *payload)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	statuses, err := s.statusService.ProjectStatuses(r.Context(), projectID)
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	web.Success(w, adapter.ToTaskStatusesDto(statuses))
}
