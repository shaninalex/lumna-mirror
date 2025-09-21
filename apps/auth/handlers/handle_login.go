// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"errors"
	"net/http"

	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
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

	user, err := s.userRepository.GetByField(ctx, db, "email", payload.Email)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password)); err != nil {
		web.Error(w, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	session, _ := s.sessionStore.Get(r, "app_session")
	session.Values["user_id"] = user.ID
	session.Values["user_email"] = user.Email
	session.Options.MaxAge = 86400 * 7
	err = s.sessionStore.Save(r, w, session)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "Login Successful")
}
