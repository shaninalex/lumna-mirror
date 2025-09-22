// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"context"
	"errors"
	"net/http"

	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models/repositories"
	"golang.org/x/crypto/bcrypt"
)

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := database.GetDb(r.Context())
	payload, err := web.BodyParser[loginPayload](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err := repositories.UserGetByField(ctx, db, "email", payload.Email)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password)); err != nil {
		web.Error(w, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	ctx = context.WithValue(ctx, "device", r.UserAgent())
	tokenString, err := s.tokenService.CreateToken(ctx, user.ID)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, map[string]any{"token": tokenString}, "Login Successful")
}
