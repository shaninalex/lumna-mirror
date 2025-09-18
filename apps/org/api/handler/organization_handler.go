// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"errors"
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/org/adapter"
	"gitlab.com/shaninalex/flowreon/apps/org/domain"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// OrganizationHandler - organization handler.
type OrganizationHandler struct {
	manager domain.OrganizationManager
}

// NewOrganizationHandler - new organization handler.
func NewOrganizationHandler(manager domain.OrganizationManager) *OrganizationHandler {
	return &OrganizationHandler{
		manager: manager,
	}
}

// HandleGetByUser - handle get by user.
// TODO: rename
func (s *OrganizationHandler) HandleGetByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	organization, err := s.manager.Get(ctx, web.GetUserID(r))
	if errors.Is(err, apperrors.OrgNotFound) {
		web.ReturnJSON(w, http.StatusNotFound, nil, err.Error())
		return
	}
	web.Success(w, adapter.ToDto(organization))
}
