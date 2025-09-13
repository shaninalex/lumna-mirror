// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package dto

import (
	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

type UserDto struct {
	ID       uuid.UUID            `json:"id"`
	Settings *models.UserSettings `json:"settings"`
	Code     string               `json:"code"`
}

func ToUserDto(user *models.User) *UserDto {
	return &UserDto{
		ID:       user.ID,
		Settings: user.GetSettings(),
		Code:     user.GetCode(),
	}
}
