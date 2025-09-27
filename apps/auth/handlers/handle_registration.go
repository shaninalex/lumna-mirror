// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"errors"
	"net/http"

	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/repositories"
	"golang.org/x/crypto/bcrypt"
)

type registerPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := database.GetDb(r.Context())
	payload, err := web.BodyParser[registerPayload](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	_, err = repositories.UserGetByField(ctx, db, "email", payload.Email)
	if err == nil {
		web.Error(w, http.StatusBadRequest, errors.New("user with email already exists"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	user := &models.User{
		Email:        payload.Email,
		PasswordHash: string(hash),
	}
	user, err = repositories.UserSave(ctx, db, user)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, nil, "Registration Successful")
}
