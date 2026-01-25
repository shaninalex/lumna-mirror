package pages

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/pkg/csrf"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web/adapters"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/board"
	boardPartials "gitlab.com/shaninalex/lumna/app/web/pages/templates/board/partials"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/components"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/projects"
	projectPartials "gitlab.com/shaninalex/lumna/app/web/pages/templates/projects/partials"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/tasks"
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
	router.POST("/projects/:projectID/edit", controller.projectEditSubmission)
	router.GET("/projects/:projectID/board/:boardID", controller.projectBoard)
	router.GET("/projects/:projectID/board/:boardID/edit", controller.projectBoardEdit)

	router.PATCH("/hx/projects/lists/:id/reorder", controller.reorderList)

	router.GET("/hx/projects/tasks/:taskID/modal", controller.taskDetail)
	router.PATCH("/hx/projects/tasks/:id/reorder", controller.reorderTask)

	router.GET("/hx/projects/board-list/modal", controller.boardListAddModal)
	router.GET("/hx/projects/board-list/:boardListID/modal", controller.boardListEditModal)
	router.POST("/hx/projects/board-list/:boardListID", controller.boardListEditSubmission)

	router.POST("/hx/projects/:projectID/board/form", controller.projectBoardAddFormSubmission)
	router.DELETE("/hx/projects/:projectID/board/:boardID", controller.projectBoardDelete)
	router.POST("/hx/projects/:projectID/board/:boardID", controller.projectBoardUpdate)

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

func (s *ProjectsController) projectEditSubmission(c *gin.Context) {
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
	utils.RenderTemplate(c, http.StatusOK, projectPartials.ProjectUpdated(payload.Title))
}

func (s *ProjectsController) projectBoard(c *gin.Context) {
	boardID := uuid.MustParse(c.Param("boardID"))
	b, err := s.projectService.BoardGet(c.Request.Context(), boardID)
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
	b, err := s.projectService.BoardGet(c.Request.Context(), boardID)
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
	_task, err := s.projectService.GetTask(c.Request.Context(), taskID)
	if err != nil {
		panic(err)
	}
	utils.RenderTemplate(c, http.StatusOK, tasks.TaskModal(&adapters.HtmlTask{Task: _task}))
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

func (s *ProjectsController) projectBoardAddFormSubmission(c *gin.Context) {
	ID := uuid.MustParse(c.Param("projectID"))
	payload := struct {
		Title string `form:"title"`
	}{}

	if err := c.ShouldBind(&payload); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, components.Alert(err.Error(), &components.AlertTypeDanger))
		return
	}

	board, err := s.projectService.BoardCreate(c.Request.Context(), ID, payload.Title)
	if err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, components.Alert(err.Error(), &components.AlertTypeDanger))
		return
	}

	c.Header("HX-Trigger", "board:created")
	utils.RenderTemplate(c, http.StatusOK, boardPartials.BoardItem(board.ID.String(), ID.String(), board.Title))
}

func (s *ProjectsController) projectBoardDelete(c *gin.Context) {
	projectID := uuid.MustParse(c.Param("projectID"))
	boardID := uuid.MustParse(c.Param("boardID"))
	if err := s.projectService.BoardDelete(c.Request.Context(), boardID); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, components.Alert(err.Error(), &components.AlertTypeDanger))
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
		utils.RenderTemplate(c, http.StatusBadRequest, components.Alert(err.Error(), &components.AlertTypeDanger))
		return
	}

	projectID := uuid.MustParse(c.Param("projectID"))
	boardID := uuid.MustParse(c.Param("boardID"))
	if err := s.projectService.BoardUpdate(c.Request.Context(), projectID, boardID, payload.Title); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, components.Alert(err.Error(), &components.AlertTypeDanger))
		return
	}
	utils.RenderTemplate(c, http.StatusOK, components.Alert("Board was updated", &components.AlertTypeSuccess))
}

func (s *ProjectsController) boardListAddModal(c *gin.Context) {
	utils.RenderTemplate(c, http.StatusOK, boardPartials.BoardListForm(
		boardPartials.BoardListFormData{
			CSRF:      csrf.GetToken(c),
			BoardList: nil,
			Url:       "",
		},
	))
}

func (s *ProjectsController) boardListEditModal(c *gin.Context) {
	boardID := uuid.MustParse(c.Param("boardListID"))
	boardList, err := s.projectService.BoardListGet(c.Request.Context(), boardID)
	if err != nil {
		panic(err)
	}
	utils.RenderTemplate(c, http.StatusOK, boardPartials.BoardListForm(
		boardPartials.BoardListFormData{
			CSRF:      csrf.GetToken(c),
			BoardList: boardList,
			Url:       fmt.Sprintf("/hx/projects/board-list/%s", boardList.ID.String()),
		},
	))
}

func (s *ProjectsController) boardListEditSubmission(c *gin.Context) {
	payload := struct {
		Title string `form:"title"`
	}{}

	if err := c.ShouldBind(&payload); err != nil {
		utils.RenderTemplate(c, http.StatusBadRequest, components.Alert(err.Error(), &components.AlertTypeDanger))
		return
	}
}
