// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/auth/domain"
	"gitlab.com/shaninalex/flowreon/apps/auth/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// AuthHooksHandler - auth hooks handler.
type AuthHooksHandler struct {
	authAPI domain.AuthHookHandler
}

// NewAuthHooksHandler - new auth hooks handler.
func NewAuthHooksHandler(api domain.AuthHookHandler) *AuthHooksHandler {
	return &AuthHooksHandler{
		authAPI: api,
	}
}

// HandleHookRegister - handle hook register.
func (s *AuthHooksHandler) HandleHookRegister(w http.ResponseWriter, r *http.Request) {
	data, err := web.BodyParser[dto.HooksKratosPayloadDTO](r)
	if err != nil {
		web.ReturnJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	err = s.authAPI.HookRegister(r.Context(), data)
	if err != nil {
		web.ReturnJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	web.Success(w, nil)
	return
}

// HandleHookVerify - handle hook verify.
func (s *AuthHooksHandler) HandleHookVerify(w http.ResponseWriter, r *http.Request) {
	web.Success(w, nil)
	return
}

// HandleHookLogin - handle hook login.
func (s *AuthHooksHandler) HandleHookLogin(w http.ResponseWriter, r *http.Request) {
	web.Success(w, nil)
	return
}
