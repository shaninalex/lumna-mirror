// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"errors"
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// ProjectHandler - project handler.
type ProjectHandler struct {
	manager domain.ProjectManager
}

// NewProjectHandler - new project handler.
func NewProjectHandler(manager domain.ProjectManager) *ProjectHandler {
	h := &ProjectHandler{
		manager: manager,
	}
	return h
}

// HandleProjectsList - handle projects list.
func (s *ProjectHandler) HandleProjectsList(w http.ResponseWriter, r *http.Request) {
	projects, err := s.manager.List(r.Context(), web.GetOrganizationID(r))
	if err != nil {
		if errors.Is(err, apperrors.ProjectNotFound) {
			web.Success(w, adapter.NewProjectsDto(nil))
			return
		}
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewProjectsDto(projects))
}

// HandleProjectTasksList - handle project tasks list.
func (s *ProjectHandler) HandleProjectTasksList(w http.ResponseWriter, r *http.Request) {
	projectCode := r.PathValue("projectCode")

	tasks, err := s.manager.TasksList(r.Context(), web.GetOrganizationID(r), projectCode)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewTasksDto(tasks))
}

// HandleProjectCreate - handle project tasks list.
func (s *ProjectHandler) HandleProjectCreate(w http.ResponseWriter, r *http.Request) {
	projectDto, err := web.BodyParser[dto.ProjectDto](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	project, err := s.manager.CreateProject(r.Context(), web.GetUserID(r), web.GetOrganizationID(r), projectDto)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewProjectDto(project))
}

// TaskFilter - task filter.
type TaskFilter struct {
	Project  string `query:"project,required"`
	TaskCode string `query:"taskCode"`
}
