// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"time"

	"github.com/shaninalex/lumna/domain"
)

type BadgeDto struct {
	ID        uint
	Title     string
	Config    domain.BadgeConfig
	CreatedAt time.Time
}

func NewBadgeDto(badge *domain.Badge) *BadgeDto {
	return &BadgeDto{
		ID:        badge.ID,
		Title:     badge.Title,
		Config:    badge.Config,
		CreatedAt: time.Now(),
	}
}

type BadgeAddToTask struct {
	BadgeId uint `json:"badgeId"`
}
