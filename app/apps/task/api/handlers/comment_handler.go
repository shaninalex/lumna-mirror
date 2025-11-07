package handlers

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/domain"
	"gitlab.com/shaninalex/lumna/app/internal/web"
)

type CommentHandler struct {
	taskManager    domain.TaskManager
	commentManager domain.CommentManager
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		taskManager:    domain.NewTaskService(),
		commentManager: domain.NewCommentService(),
	}
}

func (s *CommentHandler) Post(w http.ResponseWriter, r *http.Request) {
	comment, err := web.BodyParser[domain.Comment](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	taskId := web.UrlNumericParam(w, r, "id")
	userId := web.GetUserID(r)

	comment.TaskId = taskId
	comment.UserId = userId

	if err = s.commentManager.CreateComment(r.Context(), comment); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, comment, "New comment created")
}

func (s *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	commentId := web.UrlNumericParam(w, r, "commentId")
	if err := s.commentManager.DeleteComment(r.Context(), commentId); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "Comment deleted")
}
