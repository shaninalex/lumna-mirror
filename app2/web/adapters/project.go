package adapters

import "gitlab.com/shaninalex/lumna/app2/models"

type ProjectDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func ToProjectDto(u *models.Project) *ProjectDTO {
	return &ProjectDTO{
		Id:   u.Id,
		Name: u.Name,
	}
}

func ToProjectsDto(list []*models.Project) (projects []*ProjectDTO) {
	for _, project := range list {
		projects = append(projects, ToProjectDto(project))
	}
	return projects
}
