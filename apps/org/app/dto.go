package app

import (
	"time"

	"gitlab.com/shaninalex/jajirra/database"
)

type OrganizationDto struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToDto(o *database.Organization) *OrganizationDto {
	return &OrganizationDto{
		Title:       o.Title,
		Description: o.Description,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}
