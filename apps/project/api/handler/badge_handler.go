// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"

	"github.com/shaninalex/lumna/apps/project/adapter"
	"github.com/shaninalex/lumna/domain"
	"github.com/shaninalex/lumna/internal/web"
)

// ProjectBadgeHandler - task handler.
type ProjectBadgeHandler struct {
	badgeManager domain.BadgeManager
}

// NewProjectBadgeHandler - new task handler.
func NewProjectBadgeHandler() *ProjectBadgeHandler {
	return &ProjectBadgeHandler{
		badgeManager: domain.NewBadgeService(),
	}
}

// List - retrieve all project badges
func (s *ProjectBadgeHandler) List(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	badges, err := s.badgeManager.List(r.Context(), uint(projectID))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	result := make([]*adapter.BadgeDto, len(badges))
	for i, badge := range badges {
		result[i] = adapter.NewBadgeDto(badge)
	}
	web.Success(w, result)
}

// Create - create project badge
func (s *ProjectBadgeHandler) Create(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	payload, err := web.BodyParser[adapter.BadgeInput](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	badge := &domain.Badge{
		ProjectID: uint(projectID),
		Title:     payload.Title,
		Config:    payload.Config,
	}
	err = s.badgeManager.Create(r.Context(), badge)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewBadgeDto(badge))
}

// Delete - delete project badge
func (s *ProjectBadgeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	projectId := web.UrlNumericParam(w, r, "id")
	badgeId := web.UrlNumericParam(w, r, "badgeId")
	err := s.badgeManager.Delete(r.Context(), uint(projectId), uint(badgeId))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, "badge deleted")
}
