// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"context"
	"errors"
	"net/http"

	"gitlab.com/shaninalex/flowreon/internal/db"
	"gitlab.com/shaninalex/flowreon/internal/token"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"golang.org/x/crypto/bcrypt"
)

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	payload, err := web.BodyParser[loginPayload](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err := db.UserGetByField(ctx, db.GetDb(r.Context()), "email", payload.Email)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password)); err != nil {
		web.Error(w, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	ctx = context.WithValue(ctx, "device", r.UserAgent())
	accessToken, refreshToken, err := s.authService.Login(ctx, user.ID, r.UserAgent())
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     token.AccessTokenCookieName,
		Value:    accessToken.Token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   token.NumericAccessTokenLifeTime,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.Token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   token.NumericRefreshTokenLifeTime,
	})

	web.Success(w, nil, "Login Successful")
}
