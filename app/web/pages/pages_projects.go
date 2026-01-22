package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/projects"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type ProjectsController struct {
	projectService *services.ProjectService
}

func NewProjectsController() *ProjectsController {
	return &ProjectsController{
		projectService: services.NewProjectService(),
	}
}
func RegisterProjectPages(router *gin.RouterGroup) {
	controller := NewProjectsController()

	router.GET("/projects", controller.projectsIndex)
}

func (s *ProjectsController) projectsIndex(c *gin.Context) {
	projectsList, err := s.projectService.List(c.Request.Context())
	if err != nil {
		// TODO: error page
		panic(err)
	}
	base := utils.GetBasePage(c.Request.Context())
	base.Title = "Projects"
	pageData := projects.ProjectsPageData{
		BasePage: base,
		Projects: projectsList,
	}
	utils.RenderTemplate(c, http.StatusOK, projects.ProjectsList(pageData))
}
