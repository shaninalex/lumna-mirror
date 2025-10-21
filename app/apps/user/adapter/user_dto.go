// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"github.com/shaninalex/lumna/app/internal/db"
)

type UserDto struct {
	ID       uint             `json:"id"`
	Code     string           `json:"code"`
	Email    string           `json:"email"`
	Active   bool             `json:"active"`
	State    db.UserState     `json:"state"`
	Settings *db.UserSettings `json:"settings"`
}

func ToUserDto(user *db.User) *UserDto {
	return &UserDto{
		ID:       user.GetID(),
		Code:     user.GetCode(),
		Email:    user.Email,
		Active:   user.Active,
		State:    user.State,
		Settings: user.GetSettings(),
	}
}
