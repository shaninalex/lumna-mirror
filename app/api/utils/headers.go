package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	AccessTokenMiddlewareErrorInvalidHeader error = errors.New("invalid header")
)

func TokenFromAuthHeader(c *gin.Context) (string, bool) {
	auth := c.GetHeader("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		return "", false
	}
	return strings.TrimPrefix(auth, "Bearer "), true
}
