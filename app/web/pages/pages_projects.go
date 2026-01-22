package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	router.GET("/projects/:project_id", controller.projectDetail)
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

func (s *ProjectsController) projectDetail(c *gin.Context) {
	id := uuid.MustParse(c.Param("project_id"))
	project, err := s.projectService.Get(c.Request.Context(), id)
	if err != nil {
		// TODO: error page
		panic(err)
	}
	base := utils.GetBasePage(c.Request.Context())
	base.Title = project.Title
	pageData := projects.ProjectDetailPageData{
		BasePage: base,
		Project:  *project,
	}
	utils.RenderTemplate(c, http.StatusOK, projects.ProjectDetail(pageData))
}
