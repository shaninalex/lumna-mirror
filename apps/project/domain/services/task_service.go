// Copyright © 2025 Lumna. All rights reserved.

package services

import (
	"context"

	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
)

// TaskReader - task reader.
type TaskReader interface {
	TasksList(ctx context.Context, id string) ([]*models.Task, error)
	TaskDetail(ctx context.Context, taskCode string) (*models.Task, error)
}

// TaskWriter - task writer.
type TaskWriter interface {
	PatchTaskStatus(ctx context.Context, taskCode string, data *models.Task) error
	TaskUpdate(ctx context.Context, taskCode string, data *models.Task) error
	TaskCreate(ctx context.Context, userID uint, data *models.Task) (*models.Task, error)
}
