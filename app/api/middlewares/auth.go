package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/auth"
)

var (
	ErrorAuthMiddlewareUnauthorized = errors.New("unaythorized")
)

func AuthMiddleware(c *gin.Context) {
	accessJWTToken, err := c.Cookie("access_token")
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		return
	}

	if accessJWTToken == "" {
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		return
	}

	claims, err := auth.ParseAccessJWTToken(accessJWTToken)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		return
	}

	c.Set("userID", claims.Subject)
	c.Next()
}
