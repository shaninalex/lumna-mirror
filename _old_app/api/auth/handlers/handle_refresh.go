package handlers

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/token"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/web"
)

func (s *AuthHandler) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	refreshCookie, err := r.Cookie(token.RefreshTokenCookieName)
	if err != nil {
		http.Error(w, "missing refresh token", http.StatusUnauthorized)
		return
	}

	result, err := s.authService.RefreshAccessToken(ctx, refreshCookie.Value)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     token.AccessTokenCookieName,
		Value:    result.Token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   token.NumericAccessTokenLifeTime,
	})

	web.Success(w, nil, "Refresh Successful")
}
