package repositories

import (
	"errors"
)

var (
	ErrorNoFieldsToUpdate = errors.New("no fields to update")
	ErrorNoRowsAffected   = errors.New("no rows affected")
	ErrorInvalidPassword  = errors.New("invalid password hash")
)
