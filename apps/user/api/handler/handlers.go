// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/user/domain"
	"gitlab.com/shaninalex/flowreon/apps/user/dto"
	"gitlab.com/shaninalex/flowreon/internal/token"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
)

type UserHandler struct {
	manager     domain.UserManager
	authService token.ApiAuthService
}

// NewUserHandler - new user handler
func NewUserHandler(manager domain.UserManager) *UserHandler {
	return &UserHandler{
		manager:     manager,
		authService: token.NewAuthService(),
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

func (s *UserHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cookie, err := r.Cookie(token.RefreshTokenCookieName)
	if err != nil {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	userID := web.GetUserID(r)
	if err := s.authService.Logout(ctx, userID, cookie.Value); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	token.ClearAuthCookies(w)
	web.Success(w, nil, "Logout Successful")
}
