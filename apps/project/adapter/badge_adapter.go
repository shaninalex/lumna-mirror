// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
)

type BadgeDto struct {
	ID     uint
	Title  string
	Config *models.BadgeStatusConfig
}

func NewBadgeDto(badge *models.Badge) *BadgeDto {
	return &BadgeDto{
		ID:     badge.ID,
		Title:  badge.Title,
		Config: badge.GetConfig(),
	}
}

type BadgeInput struct {
	Title     string
	ProjectID uint
	Config    *models.BadgeStatusConfig
}
