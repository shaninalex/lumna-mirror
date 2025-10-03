// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"
	"strconv"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
	"gitlab.com/shaninalex/flowreon/apps/project/domain/services"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// ProjectHandler - project handler.
type ProjectHandler struct {
	projectService services.ProjectManager
}

// NewProjectHandler - new project handler.
func NewProjectHandler() *ProjectHandler {
	h := &ProjectHandler{
		projectService: services.NewProjectService(),
	}
	return h
}

// List - retrieve all projects
func (s *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	projects, err := s.projectService.List(r.Context())
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.ToProjectsDto(projects))
}

// Create - create a new project
func (s *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	input, err := web.BodyParser[adapter.ProjectInput](r)
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	project, err := s.projectService.CreateProject(r.Context(), &models.Project{Title: input.Title})
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	web.Success(w, adapter.ToProjectDto(project))
}

// Get - retrieve a specific project
func (s *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	project, err := s.projectService.GetProject(r.Context(), uint(projectID))
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	projectDto := adapter.ToProjectDetailDto(project)
	web.Success(w, projectDto)
}

// Delete - delete Project
func (s *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	if err := s.projectService.DeleteProject(r.Context(), uint(projectID)); err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	web.Success(w, nil, "Project deleted")
}

// Patch - update specific project
func (s *ProjectHandler) Patch(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	input, err := web.BodyParser[adapter.ProjectInput](r)
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	project, err := s.projectService.UpdateProject(r.Context(), &models.Project{
		ID:    uint(projectID),
		Title: input.Title,
	})
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, project, "Project patched")
}
