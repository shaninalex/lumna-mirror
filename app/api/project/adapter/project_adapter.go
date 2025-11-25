package adapter

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/domain"
)

type ProjectInput struct {
	Title    string `json:"title"`
	StatusId int64  `json:"status_id"`
}

// ProjectDto - project dto.
type ProjectDto struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToProjectDto - new project dto.
func ToProjectDto(p *domain.Project) *ProjectDto {
	return &ProjectDto{
		Id:        p.Id,
		Title:     p.Title,
		Code:      p.Code,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToProjectsDto - new projects dto.
func ToProjectsDto(ps []*domain.Project) []*ProjectDto {
	projects := make([]*ProjectDto, len(ps))
	for i, p := range ps {
		projects[i] = ToProjectDto(p)
	}
	return projects
}
