// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/user/domain"
	"gitlab.com/shaninalex/flowreon/apps/user/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
)

type UserHandler struct {
	manager domain.UserManager
}

// NewUserHandler - new user handler
func NewUserHandler(manager domain.UserManager) *UserHandler {
	return &UserHandler{
		manager: manager,
	}
}

// HandleGetUser - get user object
func (s *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.manager.GetUser(r.Context(), web.GetUserID(r))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, dto.ToUserDto(user))
}

// HandleUpdateSettings - update user settings
func (s *UserHandler) HandleUpdateSettings(w http.ResponseWriter, r *http.Request) {
	data, err := web.BodyParser[models.UserSettings](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	err = s.manager.UpdateUserSettings(r.Context(), web.GetUserID(r), data)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	user, err := s.manager.GetUser(r.Context(), web.GetUserID(r))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, dto.ToUserDto(user), "Settings updated")
}
