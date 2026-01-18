package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func (s *ProjectsController) GetProjects(c *gin.Context) {
	projects, err := s.projectService.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, projects)
}
