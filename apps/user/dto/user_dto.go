// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package dto

import (
	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

type UserDto struct {
	ID       uuid.UUID            `json:"id"`
	Code     string               `json:"code"`
	Email    string               `json:"email"`
	Active   bool                 `json:"active"`
	State    models.UserState     `json:"state"`
	Settings *models.UserSettings `json:"settings"`
}

func ToUserDto(user *models.User) *UserDto {
	return &UserDto{
		ID:       user.ID,
		Code:     user.GetCode(),
		Email:    user.Email,
		Active:   user.Active,
		State:    user.State,
		Settings: user.GetSettings(),
	}
}
