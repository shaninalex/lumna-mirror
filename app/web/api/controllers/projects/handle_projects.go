package projects

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/web/api/utils"
)

func (s *ProjectsController) List(c *gin.Context) {
	projects, err := s.projectService.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, projects)
}

func (s *ProjectsController) Create(c *gin.Context) {
	payload := struct {
		Title string `json:"title"`
	}{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	userID := c.Request.Context().Value(internal.ContextUserID).(uuid.UUID)

	project, err := s.projectService.Create(c.Request.Context(), payload.Title, userID)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
	}

	utils.Success(c, project)
}

func (s *ProjectsController) Patch(c *gin.Context) {
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

func (s *ProjectsController) Delete(c *gin.Context) {
	utils.Error(c, http.StatusBadRequest, fmt.Errorf("return error foregin key constraint failed"))

	// id := uuid.MustParse(c.Param("id"))
	// if err := s.projectService.Delete(c.Request.Context(), id); err != nil {
	// 	utils.Error(c, http.StatusBadRequest, err)
	// 	return
	// }
	// utils.Success(c, nil)
}
