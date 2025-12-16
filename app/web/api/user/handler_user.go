package user

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/pkg/token"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web/adapters"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: services.NewUserManager(),
		authService: services.NewAuthManager(),
	}
}

type UserHandler struct {
	userService services.UserManager
	authService services.AuthManager
}

func (s *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.userService.GetUser(r.Context(), utils.GetUserID(r))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	// IMPORTANT! we should return user through adapter
	utils.Success(w, adapters.FromUserModel(user))
}

func (s *UserHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cookie, err := r.Cookie(token.RefreshTokenCookieName)
	if err != nil {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	userID := utils.GetUserID(r)
	if err := s.authService.Logout(ctx, userID, cookie.Value); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	token.ClearAuthCookies(w)
	utils.Success(w, nil, "Logout Successful")
}
