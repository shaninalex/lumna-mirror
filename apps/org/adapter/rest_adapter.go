// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/apps/org/dto"
	"gitlab.com/shaninalex/flowreon/models"
)

func ToDto(o *models.Organization) *dto.OrganizationDto {
	return &dto.OrganizationDto{
		Title:       o.Title,
		Description: o.Description,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}
