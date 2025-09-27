// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models/repositories"
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
	projects, err := s.manager.List(r.Context())
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	pr := []*dto.ProjectDto{}
	for _, project := range projects {
		statuses, err := repositories.TaskStatusListByProject(r.Context(), database.GetDb(r.Context()), project.Code)
		if err != nil {
			web.Error(w, http.StatusInternalServerError, err)
			return
		}
		projectDto := adapter.NewProjectDto(project)
		projectDto.Statuses = adapter.NewIssueStatusesDto(statuses)
		pr = append(pr, projectDto)
	}
	web.Success(w, pr)
}

// HandleProjectCreate - handle project tasks list.
func (s *ProjectHandler) HandleProjectCreate(w http.ResponseWriter, r *http.Request) {
	projectDto, err := web.BodyParser[dto.ProjectDto](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	project, err := s.manager.CreateProject(r.Context(), web.GetUserID(r), projectDto)
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
