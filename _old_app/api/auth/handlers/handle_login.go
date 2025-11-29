package handlers

import (
	"context"
	"errors"
	"net/http"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/token"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/web"
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

	user, err := s.userService.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := s.userService.CheckPassword(ctx, user.GetID(), payload.Password); err != nil {
		web.Error(w, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	ctx = context.WithValue(ctx, "device", r.UserAgent())
	accessToken, refreshToken, err := s.authService.Login(ctx, user.Id, r.UserAgent())
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
