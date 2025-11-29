package repositories

import (
	"errors"
)

var (
	ErrorUserNoFieldsToUpdate = errors.New("no fields to update")
	ErrorUserNoRowsAffected   = errors.New("no rows affected")
	ErrorUserInvalidPassword  = errors.New("invalid password hash")
	ErrorUserUnableToUpdate   = errors.New("Unable to update user")
	ErrorUserNotFound         = errors.New("user not found")
)
