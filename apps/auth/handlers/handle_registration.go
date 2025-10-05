// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"errors"
	"net/http"

	"github.com/shaninalex/lumna/internal/db"
	"github.com/shaninalex/lumna/internal/web"
	"golang.org/x/crypto/bcrypt"
)

type registerPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	connection := db.GetDb(r.Context())
	payload, err := web.BodyParser[registerPayload](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	_, err = db.UserGetByField(ctx, connection, "email", payload.Email)
	if err == nil {
		web.Error(w, http.StatusBadRequest, errors.New("user with email already exists"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	user := &db.User{
		Email:        payload.Email,
		PasswordHash: string(hash),
	}
	user, err = db.UserSave(ctx, connection, user)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, nil, "Registration Successful")
}
