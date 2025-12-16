package project

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/web/adapters"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func (s *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	projects, err := s.projectService.List(r.Context())
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, adapters.ToProjectsDto(projects))
}

type createProjectPayload struct {
	Name string `json:"name"`
}

func (s *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	payload, err := utils.BodyParser[createProjectPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	project := &models.Project{Name: payload.Name}
	err = s.projectService.Create(r.Context(), project)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, adapters.ToProjectDto(project))
}

func (s *ProjectHandler) BoardsList(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	boards, err := s.boardService.ProjectBoards(r.Context(), uint(id))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, boards)
}

func (s *ProjectHandler) BoardsCreate(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	boards, err := s.boardService.ProjectBoards(r.Context(), uint(id))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, boards)
}
