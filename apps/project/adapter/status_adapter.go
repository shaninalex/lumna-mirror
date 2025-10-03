// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/domain"
)

// TaskStatusDto - task status dto.
type TaskStatusDto struct {
	ID       uint                     `json:"id"`
	Title    string                   `json:"title"`
	Complete bool                     `json:"complete"`
	Index    uint                     `json:"index"`
	Config   *domain.TaskStatusConfig `json:"config"`
}

// ChangeTaskStatusInput - change task status dto.
type ChangeTaskStatusInput struct {
	FromStatusID uint `json:"from_status"`
	ToStatusID   uint `json:"to_status"`
	FromIdx      uint `json:"from_idx"`
	ToIdx        uint `json:"to_idx"`
}

// NewTaskStatusDto - new issue status dto.
func NewTaskStatusDto(i domain.Status) *TaskStatusDto {
	return &TaskStatusDto{
		ID:       i.ID,
		Title:    i.Title,
		Complete: i.Completed,
		Index:    i.ListIndex,
		Config:   i.GetConfig(),
	}
}

// ToTaskStatusesDto - new tasks status dto.
func ToTaskStatusesDto(statuses []domain.Status) []*TaskStatusDto {
	dtos := make([]*TaskStatusDto, len(statuses))
	for i, status := range statuses {
		dtos[i] = NewTaskStatusDto(status)
	}
	return dtos
}
