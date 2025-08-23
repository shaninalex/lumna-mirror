package app

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
)

type ProjectDto struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	ProjectKey string    `json:"project_key"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewProjectDto(p *database.Project) *ProjectDto {
	return &ProjectDto{
		ID:         p.ID,
		Title:      p.Title,
		ProjectKey: p.ProjectKey,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func NewProjectsDto(ps []*database.Project) []*ProjectDto {
	dtos := make([]*ProjectDto, len(ps))
	for i, p := range ps {
		dtos[i] = NewProjectDto(p)
	}
	return dtos
}

type IssueDto struct {
	ID          uuid.UUID          `json:"id"`
	UserID      uuid.UUID          `json:"User_id"`
	EpicID      *uuid.UUID         `json:"Epic_id"`
	SprintID    *uuid.UUID         `json:"Sprint_id"`
	ProjectID   uuid.UUID          `json:"Project_id"`
	Assignee    string             `json:"assignee"`
	Type        database.IssueType `json:"type"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   time.Time          `json:"deleted_at"`
}

func NewIssueDto(i *database.Issue) *IssueDto {
	return &IssueDto{
		ID:          i.ID,
		UserID:      i.UserID,
		EpicID:      i.EpicID,
		SprintID:    i.SprintID,
		ProjectID:   i.ProjectID,
		Assignee:    i.Assignee,
		Type:        i.Type,
		Title:       i.Title,
		Description: i.Description,
		Status:      i.Status,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		DeletedAt:   i.DeletedAt.Time,
	}
}

func NewIssuesDto(ii []*database.Issue) []*IssueDto {
	dtos := make([]*IssueDto, len(ii))
	for i, issue := range ii {
		dtos[i] = NewIssueDto(issue)
	}
	return dtos
}
