package adapters

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
)

type UserDTO struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromUserModel(u *models.User) *UserDTO {
	return &UserDTO{
		Id:        u.Id,
		Email:     u.Email,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
