package handlers

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/apps/task/adapter"
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

func (s *CommentHandler) List(w http.ResponseWriter, r *http.Request) {
	taskID := web.UrlNumericParam(w, r, "id")
	comments, err := s.commentManager.List(r.Context(), taskID)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewCommentsDtoList(comments))
}

func (s *CommentHandler) Post(w http.ResponseWriter, r *http.Request) {
	payload, err := web.BodyParser[adapter.CommentDto](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	taskId := web.UrlNumericParam(w, r, "id")
	userId := web.GetUserID(r)

	payload.TaskId = taskId
	payload.UserId = userId

	comment := adapter.NewDomainCommentFromDto(payload)
	if err = s.commentManager.CreateComment(r.Context(), comment); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	web.Success(w, adapter.NewCommentDto(comment), "New comment created")
}

func (s *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	commentId := web.UrlNumericParam(w, r, "commentId")
	if err := s.commentManager.DeleteComment(r.Context(), commentId); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "Comment deleted")
}
