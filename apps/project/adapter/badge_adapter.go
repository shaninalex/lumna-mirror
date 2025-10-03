// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/domain"
)

type BadgeDto struct {
	ID     uint
	Title  string
	Config *domain.BadgeStatusConfig
}

func NewBadgeDto(badge *domain.Badge) *BadgeDto {
	return &BadgeDto{
		ID:     badge.ID,
		Title:  badge.Title,
		Config: badge.GetConfig(),
	}
}

type BadgeInput struct {
	Title     string
	ProjectID uint
	Config    *domain.BadgeStatusConfig
}
