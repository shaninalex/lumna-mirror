package pages

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/projects"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/projects/board"
)

type ProjectsController struct{}

func NewProjectsController() *ProjectsController {
	return &ProjectsController{}
}

func ApplyProjectRoutes(router *gin.RouterGroup) {
	ctrl := NewProjectsController()

	router.GET("/projects", ctrl.handleList)
	router.GET("/projects/:id", ctrl.handleOverview)
	router.GET("/projects/:id/edit", ctrl.handleEdit)
	router.GET("/projects/:id/board/:board_id", ctrl.handleBoardView)
	router.GET("/projects/:id/board/:board_id/edit", ctrl.handleBoardEdit)
}

func (c *ProjectsController) handleList(ctx *gin.Context) {
	component := projects.List(ctx.Request.Context())
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *ProjectsController) handleOverview(ctx *gin.Context) {
	projectID := ctx.Param("id")
	component := projects.Overview(ctx.Request.Context(), projectID)
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *ProjectsController) handleEdit(ctx *gin.Context) {
	projectID := ctx.Param("id")
	component := projects.Edit(ctx.Request.Context(), projectID)
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *ProjectsController) handleBoardView(ctx *gin.Context) {
	projectID := ctx.Param("id")
	boardID := ctx.Param("board_id")
	component := board.View(ctx.Request.Context(), projectID, boardID)
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *ProjectsController) handleBoardEdit(ctx *gin.Context) {
	projectID := ctx.Param("id")
	boardID := ctx.Param("board_id")
	component := board.Edit(ctx.Request.Context(), projectID, boardID)
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}
