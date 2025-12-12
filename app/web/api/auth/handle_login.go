package auth

import (
	"errors"
	"net/http"

	"gitlab.com/shaninalex/lumna/app/pkg/token"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	payload, err := utils.BodyParser[loginPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err := s.userService.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := s.userService.CheckPassword(ctx, user.GetId(), payload.Password); err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	accessToken, refreshToken, err := s.authService.Login(ctx, user.GetId())
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
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
		Name:     token.RefreshTokenCookieName,
		Value:    refreshToken.Token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   token.NumericRefreshTokenLifeTime,
	})

	utils.Success(w, nil, "Login Successful")
}
