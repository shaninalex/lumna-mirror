// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"errors"
	"net/http"

	"github.com/shaninalex/lumna/app/internal/web"
)

type registerPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	payload, err := web.BodyParser[registerPayload](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	if _, err = s.userService.GetUserByEmail(ctx, payload.Email); err == nil {
		web.Error(w, http.StatusBadRequest, errors.New("user with email already exists"))
		return
	}

	if _, err = s.userService.CreateUser(ctx, payload.Email, payload.Password); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, nil, "Registration Successful")
}
