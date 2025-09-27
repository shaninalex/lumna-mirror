// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"gitlab.com/shaninalex/flowreon/apps/user/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/user/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// UserController sets up HTTP routes for user-related operations.
// It registers handlers for user account management and token management.
type UserController struct {
	router *web.Router // Router used to register HTTP endpoints
}

// NewUserController creates a new instance of UserController and initializes its routes.
func NewUserController(router *web.Router) *UserController {
	c := &UserController{
		router: router,
	}
	c.init() // initialize routes
	return c
}

// init initializes the controller.
// Currently, it just calls setRoutes to register all user-related routes.
func (s *UserController) init() {
	s.setRoutes()
}

// setRoutes registers HTTP routes for user and token operations.
// It wires URL paths to the appropriate handler methods.
func (s *UserController) setRoutes() {
	// User-related endpoints
	h := handler.NewUserHandler(domain.NewUserService())
	s.router.GET("/api/v1/user/me", h.HandleGetUser)               // fetch current user info
	s.router.POST("/api/v1/user/settings", h.HandleUpdateSettings) // update user settings
	s.router.GET("/api/v1/user/logout", h.HandleLogout)            // logout user

	// Token-related endpoints
	t := handler.NewTokenHandler()
	s.router.GET("/api/v1/user/tokens", t.HandleGetUserTokens)                    // list all user tokens
	s.router.DELETE("/api/v1/user/tokens/{tokenID}", t.HandleDeleteUserToken)     // delete a specific token
	s.router.GET("/api/v1/user/tokens/{tokenID}/revoke", t.HandleRevokeUserToken) // revoke a token
}
