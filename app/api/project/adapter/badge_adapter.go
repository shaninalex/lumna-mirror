package adapter

import (
	"gitlab.com/shaninalex/lumna/app/domain"
)

type BadgeDto struct {
	Id        int64              `json:"id"`
	Title     string             `json:"title"`
	ProjectID int64              `json:"projectId"`
	Config    domain.BadgeConfig `json:"config"`
}

func NewBadgeDto(badge *domain.Badge) *BadgeDto {
	return &BadgeDto{
		Id:        badge.Id,
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
