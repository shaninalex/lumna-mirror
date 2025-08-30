// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package adapter

import (
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/models"
)

func NewProjectDto(p *models.Project) *dto.ProjectDto {
	return &dto.ProjectDto{
		ID:         p.ID,
		Title:      p.Title,
		ProjectKey: p.ProjectKey,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func NewProjectsDto(ps []*models.Project) []*dto.ProjectDto {
	dtos := make([]*dto.ProjectDto, len(ps))
	for i, p := range ps {
		dtos[i] = NewProjectDto(p)
	}
	return dtos
}

func NewTaskDto(t *models.Task) *dto.TaskDto {
	return &dto.TaskDto{
		ID:          t.ID,
		UserID:      t.UserID,
		EpicID:      t.EpicID,
		SprintID:    t.SprintID,
		ProjectID:   t.ProjectID,
		Assignee:    t.Assignee,
		Completed:   t.Completed,
		Title:       t.Title,
		Description: t.Description,
		StatusID:    t.TaskStatusID,
		ListIdx:     t.ListIndex,
		Code:        t.Code,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		DeletedAt:   &t.DeletedAt.Time,
	}
}

func NewTasksDto(ii []*models.Task) []*dto.TaskDto {
	dtos := make([]*dto.TaskDto, len(ii))
	for i, issue := range ii {
		dtos[i] = NewTaskDto(issue)
	}
	return dtos
}

func NewIssueStatusDto(i *models.TaskStatus) *dto.TaskStatusDto {
	return &dto.TaskStatusDto{
		ID:          i.GetID(),
		Title:       i.Title,
		Description: i.Description,
		Complete:    i.Complete,
		Index:       i.Index,
		Config:      i.GetConfig(),
		Tasks:       NewTasksDto(i.Tasks),
	}
}

func NewTasksStatusDto(statuses []*models.TaskStatus) []*dto.TaskStatusDto {
	dtos := make([]*dto.TaskStatusDto, len(statuses))
	for i, status := range statuses {
		dtos[i] = NewIssueStatusDto(status)
	}
	return dtos
}
