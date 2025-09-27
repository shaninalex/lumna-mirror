// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/models"
)

// NewProjectDto - new project dto.
func NewProjectDto(p *models.Project) *dto.ProjectDto {
	return &dto.ProjectDto{
		ID:        p.ID,
		Title:     p.Title,
		Code:      p.Code,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// NewProjectsDto - new projects dto.
func NewProjectsDto(ps []*models.Project) []*dto.ProjectDto {
	dtos := make([]*dto.ProjectDto, len(ps))
	for i, p := range ps {
		dtos[i] = NewProjectDto(p)
	}
	return dtos
}

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
	if t.Description != nil {
		s.Description = *t.Description
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

// NewIssueStatusDto - new issue status dto.
func NewIssueStatusDto(i *models.TaskStatus) *dto.TaskStatusDto {
	return &dto.TaskStatusDto{
		ID:       i.GetID(),
		Title:    i.Title,
		Complete: i.Completed,
		Index:    i.ListIndex,
		Config:   i.GetConfig(),
	}
}

// NewIssueStatusesDto - new tasks status dto.
func NewIssueStatusesDto(statuses []*models.TaskStatus) []*dto.TaskStatusDto {
	dtos := make([]*dto.TaskStatusDto, len(statuses))
	for i, status := range statuses {
		dtos[i] = NewIssueStatusDto(status)
	}
	return dtos
}
