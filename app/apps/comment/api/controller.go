package api

import (
	"gitlab.com/shaninalex/lumna/app/apps/comment/api/handlers"
	"gitlab.com/shaninalex/lumna/app/internal/web"
)

// CommentController - task controller.
type CommentController struct {
	router *web.Router
}

// NewCommentController - new task controller.
func NewCommentController(router *web.Router) {
	controller := CommentController{router: router}
	controller.init()
}

func (s *CommentController) init() {
	commentHandler := handlers.NewCommentHandler()
	s.router.GET("/api/v1/comments", commentHandler.List)
	s.router.POST("/api/v1/comments", commentHandler.Post)
	s.router.DELETE("/api/v1/comments/{commentId}", commentHandler.Delete)
}
