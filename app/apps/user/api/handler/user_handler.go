package handler

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/apps/user/adapter"
	"gitlab.com/shaninalex/lumna/app/domain"
	"gitlab.com/shaninalex/lumna/app/pkg/token"
	"gitlab.com/shaninalex/lumna/app/pkg/web"
)

// UserHandler handles HTTP requests related to user accounts.
// It serves as the controller layer, delegating business logic to the UserManager and AuthService.
type UserHandler struct {
	manager     domain.UserManager   // Service for user-related operations (fetching, updating, etc.)
	authService token.ApiAuthService // Service for authentication-related operations (logout, token management)
}

// NewUserHandler creates a new instance of UserHandler with the provided UserManager.
// It also initializes the authentication service.
func NewUserHandler(manager domain.UserManager) *UserHandler {
	return &UserHandler{
		manager:     manager,
		authService: token.NewAuthService(),
	}
}

// HandleGetUser handles a request to fetch the current authenticated user's information.
// Retrieves the user Id from the request context and returns a UserDto in the response.
func (s *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.manager.GetUser(r.Context(), web.GetUserID(r)) // fetch user from DB
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // return 400 if error occurs
		return
	}
	web.Success(w, adapter.ToUserDto(user)) // return user data as DTO
}

// HandleUpdateSettings handles a request to update the authenticated user's settings.
// It parses the request body into a UserSettings struct, updates the user, and returns the updated user.
func (s *UserHandler) HandleUpdateSettings(w http.ResponseWriter, r *http.Request) {
	// Parse request body into UserSettings struct
	data, err := web.BodyParser[domain.UserSettings](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // invalid body
		return
	}

	// Update user settings in the database
	err = s.manager.UpdateUserSettings(r.Context(), web.GetUserID(r), data)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // update failed
		return
	}

	// Fetch updated user and return as response
	user, err := s.manager.GetUser(r.Context(), web.GetUserID(r))
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // fetching updated user failed
		return
	}

	web.Success(w, adapter.ToUserDto(user), "Settings updated") // success response
}

// HandleLogout handles a request to log out the authenticated user.
// It reads the refresh token from cookies, invalidates it via AuthService, clears cookies, and returns a success response.
func (s *UserHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Retrieve refresh token from cookie
	cookie, err := r.Cookie(token.RefreshTokenCookieName)
	if err != nil {
		http.Error(w, "missing token", http.StatusUnauthorized) // user not logged in
		return
	}

	userID := web.GetUserID(r)

	// Call AuthService to logout (delete token)
	if err := s.authService.Logout(ctx, userID, cookie.Value); err != nil {
		web.Error(w, http.StatusBadRequest, err) // failed to delete token
		return
	}

	// Clear auth cookies from client
	token.ClearAuthCookies(w)

	web.Success(w, nil, "Logout Successful") // success response
}
