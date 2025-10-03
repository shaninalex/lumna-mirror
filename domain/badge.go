// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"encoding/json"
	"time"

	"gitlab.com/shaninalex/flowreon/internal/db"
)

type Badge struct {
	ID        uint
	ProjectID uint
	Title     string
	Config    *string
	CreatedAt time.Time
}

// SaveConfig - saves the config.
func (s *Badge) SaveConfig(cnf BadgeStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	res := string(b)
	s.Config = &res
}

// GetConfig - returns the config.
func (s *Badge) GetConfig() *BadgeStatusConfig {
	if s.Config == nil {
		return NewBadgeStatusConfig()
	}
	var config BadgeStatusConfig
	err := json.Unmarshal([]byte(*s.Config), &config)
	if err != nil {
		return NewBadgeStatusConfig()
	}
	return &config
}

// BadgeStatusConfig - badge config.
type BadgeStatusConfig struct {
	Color string `json:"color,omitempty"`
}

// NewBadgeStatusConfig - new task status config.
func NewBadgeStatusConfig() *BadgeStatusConfig {
	return &BadgeStatusConfig{
		Color: "default",
	}
}

type BadgeReader interface {
	List(ctx context.Context, projectID uint) ([]*Badge, error)
}

type BadgeWriter interface {
	Create(ctx context.Context, badge *Badge) error
	Delete(ctx context.Context, projectID, badgeID uint) error
}

type BadgeManager interface {
	BadgeReader
	BadgeWriter
}

type BadgeService struct {
}

func NewBadgeService() *BadgeService {
	return &BadgeService{}
}

func (b *BadgeService) List(ctx context.Context, projectID uint) ([]*Badge, error) {
	dbBadges, err := db.BadgeProjectList(ctx, db.GetDb(ctx), projectID)
	if err != nil {
		return nil, err
	}
	result := make([]*Badge, len(dbBadges))
	for i, badge := range dbBadges {
		result[i] = &Badge{
			ID:        badge.ID,
			ProjectID: badge.ProjectID,
			Title:     badge.Title,
			Config:    &badge.Config,
			CreatedAt: badge.CreatedAt,
		}
	}
	return nil, nil
}

func (b *BadgeService) Create(ctx context.Context, badge *Badge) error {
	dbBadge := &db.Badge{
		Title:     badge.Title,
		ProjectID: badge.ProjectID,
		Config:    *badge.Config,
		CreatedAt: time.Now(),
	}
	err := db.BadgeCreate(ctx, db.GetDb(ctx), dbBadge)
	if err != nil {
		return err
	}
	badge.ID = dbBadge.ID
	return nil
}

func (b *BadgeService) Delete(ctx context.Context, projectID, badgeID uint) error {
	return db.BadgeDelete(ctx, db.GetDb(ctx), projectID, badgeID)
}
