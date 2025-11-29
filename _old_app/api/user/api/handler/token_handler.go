package handler

import (
	"net/http"
	"strconv"

	"gitlab.com/shaninalex/lumna/_old_app/apps/user/adapter"
	"gitlab.com/shaninalex/lumna/_old_app/domain"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/web"
)

// TokenHandler handles HTTP requests related to user tokens.
// It acts as the controller layer, delegating business logic to the UserTokenManager.
type TokenHandler struct {
	tokenManager domain.UserTokenManager // Service for listing, deleting, and revoking user tokens
}

// NewTokenHandler creates a new instance of TokenHandler with a default UserTokenService.
func NewTokenHandler() *TokenHandler {
	return &TokenHandler{
		tokenManager: domain.NewUserTokenService(), // inject concrete service implementation
	}
}

// HandleGetUserTokens handles a request to list all tokens for the authenticated user.
// It retrieves the user Id from the request context, fetches tokens via the service,
// converts them to DTOs, and returns them as JSON.
func (s *TokenHandler) HandleGetUserTokens(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tokens, err := s.tokenManager.List(ctx, web.GetUserID(r)) // fetch all tokens for the user
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // return 400 if any error occurs
		return
	}
	web.Success(w, adapter.ToUserTokenDtoList(tokens)) // return list of tokens as DTOs
}

// HandleDeleteUserToken handles a request to delete a specific token for the authenticated user.
// It reads the token Id from the URL, calls the service to delete the token, and responds with success/error.
func (s *TokenHandler) HandleDeleteUserToken(w http.ResponseWriter, r *http.Request) {
	tokenID, err := strconv.ParseInt(r.PathValue("tokenID"), 10, 64) // parse token Id from URL
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // invalid token Id format
		return
	}
	err = s.tokenManager.Delete(r.Context(), web.GetUserID(r), tokenID) // delete token
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // deletion failed
		return
	}
	web.Success(w, nil, "Token removed") // return success response
}

// HandleRevokeUserToken handles a request to revoke a specific token for the authenticated user.
// Revoking a token marks it as invalid without deleting it, preventing further use.
func (s *TokenHandler) HandleRevokeUserToken(w http.ResponseWriter, r *http.Request) {
	tokenID, err := strconv.ParseInt(r.PathValue("tokenID"), 10, 64) // parse token Id from URL
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // invalid token Id format
		return
	}
	err = s.tokenManager.Revoke(r.Context(), web.GetUserID(r), tokenID) // revoke token
	if err != nil {
		web.Error(w, http.StatusBadRequest, err) // revocation failed
		return
	}
	web.Success(w, nil, "Token revoked") // return success response
}
