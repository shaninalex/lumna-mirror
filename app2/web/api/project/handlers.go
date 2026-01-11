package project

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app2/models"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
	"gitlab.com/shaninalex/lumna/app2/web/adapters"
	"gitlab.com/shaninalex/lumna/app2/web/utils"
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

func (s *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	err := s.projectService.Delete(r.Context(), uint(id))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, nil, "Project deleted")
}

func (s *ProjectHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	payload, err := utils.BodyParser[createProjectPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	err = s.projectService.Patch(r.Context(), uint(id), db.Set(
		db.Field("name", payload.Name),
	))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	project, err := s.projectService.Get(r.Context(), uint(id))
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

type createBoardPayload struct {
	Name string `json:"name"`
}

func (s *ProjectHandler) BoardsCreate(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	payload, err := utils.BodyParser[createBoardPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	board, err := s.boardService.Create(r.Context(), &models.Board{
		Name:      payload.Name,
		ProjectId: uint(id),
	})
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, board)
}
