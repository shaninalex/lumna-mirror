// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package dto

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

type ProjectDto struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	ProjectKey string    `json:"project_key"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewProjectDto(p *models.Project) *ProjectDto {
	return &ProjectDto{
		ID:         p.ID,
		Title:      p.Title,
		ProjectKey: p.ProjectKey,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func NewProjectsDto(ps []*models.Project) []*ProjectDto {
	dtos := make([]*ProjectDto, len(ps))
	for i, p := range ps {
		dtos[i] = NewProjectDto(p)
	}
	return dtos
}

type TaskDto struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"creator_id"`
	EpicID      *uuid.UUID `json:"epic_id"`
	SprintID    *uuid.UUID `json:"sprint_id"`
	ProjectID   uuid.UUID  `json:"project_id"`
	Assignee    string     `json:"assignee"`
	Completed   bool       `json:"completed"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StatusID    uuid.UUID  `json:"status"`
	ListIdx     uint       `json:"list_idx"`
	Code        string     `json:"code"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func NewTaskDto(t *models.Task) *TaskDto {
	return &TaskDto{
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

func NewTasksDto(ii []*models.Task) []*TaskDto {
	dtos := make([]*TaskDto, len(ii))
	for i, issue := range ii {
		dtos[i] = NewTaskDto(issue)
	}
	return dtos
}

type TaskStatusDto struct {
	ID          uuid.UUID                `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Complete    bool                     `json:"complete"`
	Index       uint                     `json:"index"`
	Config      *models.TaskStatusConfig `json:"config"`
	Tasks       []*TaskDto               `json:"tasks"`
}

func NewIssueStatusDto(i *models.TaskStatus) *TaskStatusDto {
	return &TaskStatusDto{
		ID:          i.GetID(),
		Title:       i.Title,
		Description: i.Description,
		Complete:    i.Complete,
		Index:       i.Index,
		Config:      i.GetConfig(),
		Tasks:       NewTasksDto(i.Tasks),
	}
}

func NewTasksStatusDto(statuses []*models.TaskStatus) []*TaskStatusDto {
	dtos := make([]*TaskStatusDto, len(statuses))
	for i, status := range statuses {
		dtos[i] = NewIssueStatusDto(status)
	}
	return dtos
}
