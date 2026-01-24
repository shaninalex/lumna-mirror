package pages

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/partials"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/projects"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/projects/board"
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
	router.GET("/projects/:projectID", controller.projectDetail)
	router.GET("/projects/:projectID/board/:boardID", controller.projectBoard)
	router.GET("/projects/:projectID/board/:boardID/edit", controller.projectBoardEdit)
	router.GET("/hx/tasks/:taskID/modal", controller.taskDetail)

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
	ID := uuid.MustParse(c.Param("projectID"))
	project, err := s.projectService.Get(c.Request.Context(), ID)
	if err != nil {
		panic(err)
	}
	base := utils.GetBasePage(c.Request.Context())
	base.Title = fmt.Sprintf("Project: %s", project.Title)
	pageData := projects.ProjectDetailPageData{
		BasePage: base,
		Project:  *project,
	}
	utils.RenderTemplate(c, http.StatusOK, projects.ProjectDetail(pageData))
}

func (s *ProjectsController) projectBoard(c *gin.Context) {
	boardID := uuid.MustParse(c.Param("boardID"))
	base := utils.GetBasePage(c.Request.Context())
	b, err := s.projectService.GetBoard(c.Request.Context(), boardID)
	if err != nil {
		panic(err)
	}
	base.Title = b.Title
	pageData := board.BoardPageData{
		BasePage: base,
		Board:    *b,
	}
	utils.RenderTemplate(c, http.StatusOK, board.BoardDetail(pageData))
}

func (s *ProjectsController) projectBoardEdit(c *gin.Context) {
	boardID := uuid.MustParse(c.Param("boardID"))
	base := utils.GetBasePage(c.Request.Context())
	b, err := s.projectService.GetBoard(c.Request.Context(), boardID)
	if err != nil {
		panic(err)
	}
	base.Title = b.Title
	pageData := board.BoardPageData{
		BasePage: base,
		Board:    *b,
	}
	utils.RenderTemplate(c, http.StatusOK, board.BoardEdit(pageData))
}

func (s *ProjectsController) taskDetail(c *gin.Context) {
	taskID := uuid.MustParse(c.Param("taskID"))
	task, err := s.projectService.GetTask(c.Request.Context(), taskID)
	if err != nil {
		panic(err)
	}
	utils.RenderTemplate(c, http.StatusOK, partials.TaskModal(task))
}
