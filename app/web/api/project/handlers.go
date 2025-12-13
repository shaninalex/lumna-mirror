package project

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func (s *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	projects := []*models.Project{}
	utils.Success(w, projects)
}

func (s *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	projects := []*models.Project{}
	utils.Success(w, projects)
}
