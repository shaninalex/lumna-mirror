package handlers

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/apps/task/adapter"
	"gitlab.com/shaninalex/lumna/app/domain"
	"gitlab.com/shaninalex/lumna/app/internal/web"
)

type BadgeHandler struct {
	badgeWriter domain.BadgeWriter
}

func NewBadgeHandler() *BadgeHandler {
	return &BadgeHandler{}
}

func (s *BadgeHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID := web.UrlNumericParam(w, r, "id")
	payload, err := web.BodyParser[adapter.BadgeAddToTask](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = s.badgeWriter.AddToTask(ctx, taskID, payload.BadgeId); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "badge added to task")
}

func (s *BadgeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID := web.UrlNumericParam(w, r, "id")
	badgeID := web.UrlNumericParam(w, r, "badgeId")
	if err := s.badgeWriter.DeleteFromTask(ctx, taskID, badgeID); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "badge added to task")
}
