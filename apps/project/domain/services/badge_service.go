// Copyright © 2025 Lumna. All rights reserved.

package services

import (
	"context"
	"time"

	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
	"gitlab.com/shaninalex/flowreon/internal/db"
)

type BadgeProjectReader interface {
	List(ctx context.Context, projectID uint) ([]*models.Badge, error)
}

type BadgeProjectWriter interface {
	Create(ctx context.Context, badge *models.Badge) error
	Delete(ctx context.Context, projectID, badgeID uint) error
}

type BadgeProjectManager interface {
	BadgeProjectReader
	BadgeProjectWriter
}

type BadgeProjectService struct {
}

func NewBadgeProjectService() *BadgeProjectService {
	return &BadgeProjectService{}
}

func (b *BadgeProjectService) List(ctx context.Context, projectID uint) ([]*models.Badge, error) {
	dbBadges, err := db.BadgeProjectList(ctx, db.GetDb(ctx), projectID)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Badge, len(dbBadges))
	for i, badge := range dbBadges {
		result[i] = &models.Badge{
			ID:        badge.ID,
			ProjectID: badge.ProjectID,
			Title:     badge.Title,
			Config:    &badge.Config,
			CreatedAt: badge.CreatedAt,
		}
	}
	return nil, nil
}

func (b *BadgeProjectService) Create(ctx context.Context, badge *models.Badge) error {
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

func (b *BadgeProjectService) Delete(ctx context.Context, projectID, badgeID uint) error {
	return db.BadgeDelete(ctx, db.GetDb(ctx), projectID, badgeID)
}
