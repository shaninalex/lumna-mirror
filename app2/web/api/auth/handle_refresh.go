package auth

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app2/pkg/token"
	"gitlab.com/shaninalex/lumna/app2/web/utils"
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
		utils.Error(w, http.StatusBadRequest, err)
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

	utils.Success(w, nil, "Refresh Successful")
}
