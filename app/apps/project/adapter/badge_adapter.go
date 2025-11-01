// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/lumna/app/domain"
)

type BadgeDto struct {
	ID        int64              `json:"id"`
	Title     string             `json:"title"`
	ProjectID int64              `json:"projectId"`
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
	ProjectID int64              `json:"projectId"`
	Config    domain.BadgeConfig `json:"config"`
}
