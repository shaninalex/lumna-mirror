// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"
	"strconv"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// ProjectHandler - project handler.
type ProjectHandler struct {
	projectReader domain.ProjectReader
	projectWriter domain.ProjectWriter
}

// NewProjectHandler - new project handler.
func NewProjectHandler() *ProjectHandler {
	h := &ProjectHandler{}
	return h
}

// List - retrieve all projects
func (s *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	projects, err := s.projectReader.List(r.Context())
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, http.StatusOK, adapter.ToProjectsDto(projects))
}

// Create - create a new project
func (s *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {}

// Get - retrieve a specific project
func (s *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	tokenID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	project, err := s.projectReader.GetProject(r.Context(), uint(tokenID))
	if err != nil {
		web.Error(w, http.StatusNotFound, err)
		return
	}
	projectDto := adapter.ToProjectDetailDto(project)
	web.Success(w, http.StatusOK, projectDto)
}

// Delete - delete Project
func (s *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {}

// Patch - update specific project
func (s *ProjectHandler) Patch(w http.ResponseWriter, r *http.Request) {
	// This request return project WITH columns. Returns after update
}
