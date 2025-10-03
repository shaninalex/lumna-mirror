// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"time"

	"gitlab.com/shaninalex/flowreon/domain"
)

type BadgeDto struct {
	ID        uint
	Title     string
	Config    *domain.BadgeStatusConfig
	CreatedAt time.Time
}

func NewBadgeDto(badge *domain.Badge) *BadgeDto {
	return &BadgeDto{
		ID:        badge.ID,
		Title:     badge.Title,
		Config:    badge.GetConfig(),
		CreatedAt: time.Now(),
	}
}
