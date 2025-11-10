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
	userId := web.GetUserID(r)
	comment.UserId = userId

	if err = s.commentManager.CreateComment(r.Context(), comment); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, comment, "New comment created")
}

func (s *CommentHandler) List(w http.ResponseWriter, r *http.Request) {
	entityId := web.UrlNumericQueryParam(w, r, "entity_id")
	entityType := r.URL.Query().Get("entity_type")
	comments := s.commentManager.ListComments(r.Context(), entityId, entityType)
	web.Success(w, comments)
}

func (s *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	commentId := web.UrlNumericParam(w, r, "commentId")
	if err := s.commentManager.DeleteComment(r.Context(), commentId); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "Comment deleted")
}
