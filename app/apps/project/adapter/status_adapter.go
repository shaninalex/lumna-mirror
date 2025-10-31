// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"github.com/shaninalex/lumna/app/domain"
)

// TaskStatusDto - task status dto.
type TaskStatusDto struct {
	ID        int64                    `json:"id"`
	Title     string                   `json:"title"`
	ProjectId int64                    `json:"project_id"`
	Index     int64                    `json:"index"`
	Config    *domain.TaskStatusConfig `json:"config"`
}

// NewTaskStatusDto - new issue status dto.
func NewTaskStatusDto(i *domain.Status) *TaskStatusDto {
	return &TaskStatusDto{
		ID:        i.ID,
		Title:     i.Title,
		Index:     i.ListIndex,
		ProjectId: i.ProjectId,
		Config:    i.GetConfig(),
	}
}

// ToTaskStatusesDto - new tasks status dto.
func ToTaskStatusesDto(statuses []*domain.Status) []*TaskStatusDto {
	dtos := make([]*TaskStatusDto, len(statuses))
	for i, status := range statuses {
		dtos[i] = NewTaskStatusDto(status)
	}
	return dtos
}

type TaskStatusInput struct {
	Title string `json:"title"`
}
