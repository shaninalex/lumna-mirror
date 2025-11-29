package adapter

import (
	"gitlab.com/shaninalex/lumna/_old_app/domain"
)

type UserDto struct {
	Id       int64                `json:"id"`
	Code     string               `json:"code"`
	Email    string               `json:"email"`
	Active   bool                 `json:"active"`
	State    domain.UserState     `json:"state"`
	Settings *domain.UserSettings `json:"settings"`
}

func ToUserDto(user *domain.User) *UserDto {
	return &UserDto{
		Id:       user.GetID(),
		Code:     user.GetCode(),
		Email:    user.Email,
		Active:   user.Active,
		State:    user.State,
		Settings: user.GetSettings(),
	}
}
