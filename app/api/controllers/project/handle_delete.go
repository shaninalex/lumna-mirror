package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ProjectController) Delete(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	if err := s.projectService.Delete(c.Request.Context(), id); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, nil)
}
