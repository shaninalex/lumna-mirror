// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/domain"
)

type BadgeDto struct {
	ID     uint
	Title  string
	Config domain.BadgeConfig
}

func NewBadgeDto(badge *domain.Badge) *BadgeDto {
	return &BadgeDto{
		ID:     badge.ID,
		Title:  badge.Title,
		Config: badge.Config,
	}
}

type BadgeInput struct {
	Title     string
	ProjectID uint
	Config    domain.BadgeConfig
}
