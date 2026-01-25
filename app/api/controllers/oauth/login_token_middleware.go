package oauth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/jwt"
)

func OAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := jwt.ValidateLoginToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.Set("userID", userID)
	c.Next()
}
