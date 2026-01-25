package pages

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/pkg/csrf"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web/adapters"
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
	router.GET("/projects/:projectID/edit", controller.projectEdit)
	router.POST("/projects/:projectID/edit", controller.projectEditSubmit)
	router.GET("/projects/:projectID/board/:boardID", controller.projectBoard)
	router.GET("/projects/:projectID/board/:boardID/edit", controller.projectBoardEdit)

	router.GET("/hx/tasks/:taskID/modal", controller.taskDetail)
	router.POST("/hx/projects/:projectID/board/form", controller.projectBoardAddFormSubmition)
	router.PATCH("/hx/projects/tasks/:id/reorder", controller.reorderTask)
	router.PATCH("/hx/projects/lists/:id/reorder", controller.reorderList)
	router.DELETE("/hx/projects/:projectID/board/:boardID", controller.projectBoardDelete)
	router.POST("/hx/projects/:projectID/board/:boardID", controller.projectBoardUpdate)

	router.GET("/hx/projects/board-list/modal", controller.boardListModal)
}

func (s *ProjectsController) projectsIndex(c *gin.Context) {
	projectsList, err := s.projectService.List(c.Request.Context())
	if err != nil {
		// TODO: error page
		panic(err)
	}
	pageData := projects.ProjectsPageData{
		BasePage: utils.BasePageData(c, "Projects"),
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
	pageData := projects.ProjectDetailPageData{
		BasePage: utils.BasePageData(c, fmt.Sprintf("Project: %s", project.Title)),
		Project:  *project,
	}
	utils.RenderTemplate(c, http.StatusOK, projects.ProjectDetail(pageData))
}

func (s *ProjectsController) projectEdit(c *gin.Context) {
	ID := uuid.MustParse(c.Param("projectID"))
	project, err := s.projectService.Get(c.Request.Context(), ID)
	if err != nil {
		panic(err)
	}
	pageData := projects.ProjectDetailPageData{
		BasePage: utils.BasePageData(c, fmt.Sprintf("Project: %s", project.Title)),
		Project:  *project,
	}
	utils.RenderTemplate(c, http.StatusOK, projects.ProjectEdit(pageData))
}

func (s *ProjectsController) projectEditSubmit(c *gin.Context) {
	ID := uuid.MustParse(c.Param("projectID"))
	project, err := s.projectService.Get(c.Request.Context(), ID)
	if err != nil {
		panic(err)
	}
	payload := struct {
		Title string `form:"title"`
	}{}
	if err := c.ShouldBind(&payload); err != nil {
		panic(err)
	}
	if err := s.projectService.UpdateProject(c.Request.Context(), project.ID, payload.Title); err != nil {
		panic(err)
	}
	utils.RenderTemplate(c, http.StatusOK, partials.ProjectUpdated(payload.Title))
}

func (s *ProjectsController) projectBoard(c *gin.Context) {
	boardID := uuid.MustParse(c.Param("boardID"))
	b, err := s.projectService.GetBoard(c.Request.Context(), boardID)
	if err != nil {
		panic(err)
	}
	pageData := board.BoardPageData{
		BasePage: utils.BasePageData(c, fmt.Sprintf("Board: %s", b.Title)),
		Board:    *b,
	}
	utils.RenderTemplate(c, http.StatusOK, board.BoardDetail(pageData))
}

func (s *ProjectsController) projectBoardEdit(c *gin.Context) {
	boardID := uuid.MustParse(c.Param("boardID"))
	b, err := s.projectService.GetBoard(c.Request.Context(), boardID)
	if err != nil {
		panic(err)
	}
	pageData := board.BoardPageData{
		BasePage: utils.BasePageData(c, fmt.Sprintf("Edit board: %s", b.Title)),
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
	utils.RenderTemplate(c, http.StatusOK, partials.TaskModal(&adapters.HtmlTask{Task: task}))
}

func (s *ProjectsController) reorderTask(c *gin.Context) {
	payload := struct {
		BoardListID string `json:"boardListId"`
		Order       uint   `json:"order"`
	}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	taskID := uuid.MustParse(c.Param("id"))
	boardListID := uuid.MustParse(payload.BoardListID)

	if err := s.projectService.ReorderTask(c.Request.Context(), taskID, boardListID, payload.Order); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, nil)
}

func (s *ProjectsController) reorderList(c *gin.Context) {
	payload := struct {
		Order uint `json:"order"`
	}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	listID := uuid.MustParse(c.Param("id"))

	if err := s.projectService.ReorderList(c.Request.Context(), listID, payload.Order); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, nil)
}

func (s *ProjectsController) projectBoardAddFormSubmition(c *gin.Context) {
	ID := uuid.MustParse(c.Param("projectID"))
	payload := struct {
		Title string `form:"title"`
	}{}

	if err := c.ShouldBind(&payload); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}

	board, err := s.projectService.BoardCreate(c.Request.Context(), ID, payload.Title)
	if err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}

	c.Header("HX-Trigger", "board:created")
	utils.RenderTemplate(c, http.StatusOK, partials.BoardItem(board.ID.String(), ID.String(), board.Title))
}

func (s *ProjectsController) projectBoardDelete(c *gin.Context) {
	projectID := uuid.MustParse(c.Param("projectID"))
	boardID := uuid.MustParse(c.Param("boardID"))
	if err := s.projectService.BoardDelete(c.Request.Context(), boardID); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}
	redirectURL := fmt.Sprintf("/projects/%s", projectID)
	c.Header("HX-Trigger", "board:deleted")
	c.Writer.Header().Set("HX-Redirect", redirectURL)
}

func (s *ProjectsController) projectBoardUpdate(c *gin.Context) {
	payload := struct {
		Title string `form:"title"`
	}{}

	if err := c.ShouldBind(&payload); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}

	projectID := uuid.MustParse(c.Param("projectID"))
	boardID := uuid.MustParse(c.Param("boardID"))
	if err := s.projectService.BoardUpdate(c.Request.Context(), projectID, boardID, payload.Title); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}
	utils.RenderTemplate(c, http.StatusOK, partials.Alert("Board was updated", &partials.AlertTypeSuccess))
}

func (s *ProjectsController) boardListModal(c *gin.Context) {
	utils.RenderTemplate(c, http.StatusOK, partials.BoardListForm(
		csrf.GetToken(c),
		false,
	))
}
