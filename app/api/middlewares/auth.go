package middlewares

import (
	"errors"
	"log"
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
		log.Println(1)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		c.Abort()
		return
	}

	if accessJWTToken == "" {
		log.Println(2)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		c.Abort()
		return
	}

	claims, err := auth.ParseAccessJWTToken(accessJWTToken)
	if err != nil {
		log.Println(3)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
		c.Abort()
		return
	}

	c.Set("userID", claims.Subject)
	c.Next()
}
