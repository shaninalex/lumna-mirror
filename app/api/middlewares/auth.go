package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/auth"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	token, ok := utils.TokenFromAuthHeader(c)
	if !ok {
		utils.Error(c, http.StatusBadRequest, utils.AccessTokenMiddlewareErrorInvalidHeader)
		return
	}
	claims, err := auth.ParseAccessJWTToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.Set("userID", claims.Subject)
	c.Next()
}
