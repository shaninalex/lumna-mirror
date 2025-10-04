// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/domain"
)

type BadgeDto struct {
	ID        uint               `json:"id"`
	Title     string             `json:"title"`
	ProjectID uint               `json:"projectId"`
	Config    domain.BadgeConfig `json:"config"`
}

func NewBadgeDto(badge *domain.Badge) *BadgeDto {
	return &BadgeDto{
		ID:        badge.ID,
		Title:     badge.Title,
		ProjectID: badge.ProjectID,
		Config:    badge.Config,
	}
}

type BadgeInput struct {
	Title     string             `json:"title"`
	ProjectID uint               `json:"projectId"`
	Config    domain.BadgeConfig `json:"config"`
}
