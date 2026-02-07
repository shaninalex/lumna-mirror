package project

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ProjectController) Create(c *gin.Context) {
	payload := struct {
		Title string `json:"title"`
	}{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
	}

	project, err := s.projectService.Create(c.Request.Context(), payload.Title, userID)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
	}

	s.logger.Log(fmt.Sprintf("Project created: %s", project.Title))

	utils.Success(c, project)
}
