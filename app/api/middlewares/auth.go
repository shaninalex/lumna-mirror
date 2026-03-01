package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/auth"
)

var (
	ErrorAuthMiddlewareUnauthorized = errors.New("unauthorized")
	ErrorAuthMiddlewareInvalidID    = errors.New("invalid user id")
)

func AuthMiddleware(c *gin.Context) {
	accessJWTToken, err := c.Cookie("access_token")
	if err != nil {
		log.Printf("[AuthMiddleware]: %s", err)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		c.Abort()
		return
	}

	if accessJWTToken == "" {
		log.Printf("[AuthMiddleware]: access token \"%s\" is invalid", accessJWTToken)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		c.Abort()
		return
	}

	claims, err := auth.ParseAccessJWTToken(accessJWTToken)
	if err != nil {
		log.Printf("[AuthMiddleware]: unable to parse: %s", err)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		c.Abort()
		return
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareInvalidID)
		c.Abort()
		return
	}

	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, internal.ContextIdentity, uint(userID))
	// fill request context with values here

	c.Set("userID", uint(userID))
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
