// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"
	"database/sql"
	"net/http"

	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
)

// AuthTestRouter - auth router.
func AuthTestRouter(ctx context.Context) *web.Router {
	db := ctx.Value(base.ContextDB).(*sql.DB)
	router := web.DefaultRouter(db, "test_router")
	kratosClient := NewMockKratosService()
	router.Use(web.NewAuthMiddleware(kratosClient).Wrap)
	return router
}

// TestRouter - basic router
func TestRouter(ctx context.Context) *web.Router {
	db := ctx.Value(base.ContextDB).(*sql.DB)
	router := web.DefaultRouter(db, "test_router")
	return router
}

// SetAuthRequest - sets the auth request.
func SetAuthRequest(req *http.Request, user models.AuthUser, cookie string) {
	req.Header.Set("X-USER", user.GetID().String())
	req.Header.Set("Cookie", cookie)
}
