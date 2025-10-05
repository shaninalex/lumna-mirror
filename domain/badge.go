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
	Config    BadgeConfig
	CreatedAt time.Time
}

// BadgeConfig - badge config.
type BadgeConfig struct {
	Color string `json:"color,omitempty"`
}

// ToBadgeConfig - converts string to badge config.
func ToBadgeConfig(cnf string) BadgeConfig {
	var config BadgeConfig
	err := json.Unmarshal([]byte(cnf), &config)
	if err != nil {
		return NewBadgeStatusConfig()
	}
	return config
}

// NewBadgeStatusConfig - new task status config.
func NewBadgeStatusConfig() BadgeConfig {
	return BadgeConfig{
		Color: "default",
	}
}

type BadgeReader interface {
	List(ctx context.Context, projectID uint) ([]*Badge, error)
}

type BadgeWriter interface {
	Create(ctx context.Context, badge *Badge) error
	Delete(ctx context.Context, projectID, badgeID uint) error
	AddToTask(ctx context.Context, taskID, badgeID uint) error
	DeleteFromTask(ctx context.Context, taskID, badgeID uint) error
}

type BadgeManager interface {
	BadgeReader
	BadgeWriter
}

type BadgeService struct {
}

func (b *BadgeService) AddToTask(ctx context.Context, taskID, badgeID uint) error {
	//TODO implement me
	panic("implement me")
}

func (b *BadgeService) DeleteFromTask(ctx context.Context, taskID, badgeID uint) error {
	//TODO implement me
	panic("implement me")
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
			Config:    ToBadgeConfig(badge.Config),
			CreatedAt: badge.CreatedAt,
		}
	}
	return nil, nil
}

func (b *BadgeService) Create(ctx context.Context, badge *Badge) error {
	cnf, err := json.Marshal(badge.Config)
	if err != nil {
		return err
	}
	dbBadge := &db.Badge{
		Title:     badge.Title,
		ProjectID: badge.ProjectID,
		Config:    string(cnf),
		CreatedAt: time.Now(),
	}
	err = db.BadgeCreate(ctx, db.GetDb(ctx), dbBadge)
	if err != nil {
		return err
	}
	badge.ID = dbBadge.ID
	return nil
}

func (b *BadgeService) Delete(ctx context.Context, projectID, badgeID uint) error {
	return db.BadgeDelete(ctx, db.GetDb(ctx), projectID, badgeID)
}
