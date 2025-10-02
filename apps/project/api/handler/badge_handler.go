// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"
)

// ProjectBadgeHandler - task handler.
type ProjectBadgeHandler struct {
}

// NewProjectBadgeHandler - new task handler.
func NewProjectBadgeHandler() *ProjectBadgeHandler {
	return &ProjectBadgeHandler{}
}

// List - retrieve all project badges
func (s *ProjectBadgeHandler) List(w http.ResponseWriter, r *http.Request) {}

// Create - create project badge
func (s *ProjectBadgeHandler) Create(w http.ResponseWriter, r *http.Request) {}

// Delete - delete project badge
func (s *ProjectBadgeHandler) Delete(w http.ResponseWriter, r *http.Request) {}
