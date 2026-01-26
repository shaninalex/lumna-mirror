package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var UserIDNotInContextError error = errors.New("user id not found in request context")

func GetUserID(c *gin.Context) (uuid.UUID, error) {
	userIDAny, ok := c.Get("userID")
	if !ok {
		return uuid.Nil, UserIDNotInContextError
	}

	userID := uuid.MustParse(userIDAny.(string))
	return userID, nil
}
