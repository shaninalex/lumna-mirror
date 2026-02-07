package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ProjectController) Patch(c *gin.Context) {
	payload := struct {
		Title string `json:"title"`
	}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	id := uuid.MustParse(c.Param("id"))
	project, err := s.projectService.Update(c.Request.Context(), id, payload.Title)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, project)
}
