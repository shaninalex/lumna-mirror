package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services"
)

type GenerateKeyPayload struct {
	Name string `json:"name"`
}

type GenerateKeyResponse struct {
	ProjectKey string `json:"project_key"`
}

func (s *ProjectController) GenerateKey(c *gin.Context) {
	var payload GenerateKeyPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	key, err := s.projectService.GenerateKey(c.Request.Context(), payload.Name)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	if key == "" {
		utils.Error(c, http.StatusNotFound, services.ErrorProjectUnableGenerateKey)
		return
	}

	utils.Success(c, &GenerateKeyResponse{
		ProjectKey: key,
	})
}
