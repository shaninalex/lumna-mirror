// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/models"
)

// NewTaskDto - new task dto.
func NewTaskDto(t *models.Task) *dto.TaskDto {
	s := &dto.TaskDto{
		ID:        t.ID,
		UserID:    t.UserID,
		ProjectID: t.ProjectID,
		Completed: t.Completed,
		Title:     t.Title,
		StatusID:  t.StatusID,
		ListIdx:   t.ListIndex,
		Code:      t.Code,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
	return s
}

// NewTasksDto - new tasks dto.
func NewTasksDto(ii []*models.Task) []*dto.TaskDto {
	dtos := make([]*dto.TaskDto, len(ii))
	for i, issue := range ii {
		dtos[i] = NewTaskDto(issue)
	}
	return dtos
}
