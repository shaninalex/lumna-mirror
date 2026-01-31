package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ErrorUserIDNotInContext error = errors.New("user id not found in request context")
var ErrorInvalidUserID error = errors.New("invalid user id")

func GetUserID(c *gin.Context) (uuid.UUID, error) {
	userIDAny, ok := c.Get("userID")
	if !ok {
		return uuid.Nil, ErrorUserIDNotInContext
	}

	userID, ok := userIDAny.(uuid.UUID)
	if !ok {
		return uuid.Nil, ErrorInvalidUserID
	}
	return userID, nil
}
