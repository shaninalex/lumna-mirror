package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserIDNotInContext error = errors.New("user id not found in request context")
var ErrorInvalidUserID error = errors.New("invalid user id")

func GetUserID(c *gin.Context) (uint, error) {
	userIDAny, ok := c.Get("userID")
	if !ok {
		return 0, ErrorUserIDNotInContext
	}

	userID, ok := userIDAny.(uint)
	if !ok {
		return 0, ErrorInvalidUserID
	}
	return userID, nil
}
