// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"
)

// ProjectTaskHandler - task handler.
type ProjectTaskHandler struct {
}

// NewProjectTaskHandler - new task handler.
func NewProjectTaskHandler() *ProjectTaskHandler {
	return &ProjectTaskHandler{}
}

// List - retrieve tasks for a project
func (s *ProjectTaskHandler) List(w http.ResponseWriter, r *http.Request) {}

// Create - create a new task in a project
func (s *ProjectTaskHandler) Create(w http.ResponseWriter, r *http.Request) {}
