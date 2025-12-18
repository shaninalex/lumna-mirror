package repositories

import (
	"errors"
)

var (
	ErrorNoRowsAffected    = errors.New("no rows affected")
	ErrorNoFieldsForUpdate = errors.New("no fields to update")

	ErrorProjectNotFound = errors.New("project not found")

	ErrorUserNoFieldsToUpdate = errors.New("no fields to update")
	ErrorUserInvalidPassword  = errors.New("invalid password hash")
	ErrorUserUnableToUpdate   = errors.New("unable to update user")
	ErrorUserNotFound         = errors.New("user not found")
)
