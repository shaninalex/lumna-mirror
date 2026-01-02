package task

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func (s *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	taskId := utils.UrlNumericParam(w, r, "id")
	task, err := s.tasksService.Get(r.Context(), uint(taskId))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, task)
}

type taskPayloadModel struct {
	Name   string `json:"name"`
	ListId uint   `json:"list_id"`
}

func (s *TaskHandler) Patch(w http.ResponseWriter, r *http.Request) {
	taskId := utils.UrlNumericParam(w, r, "id")
	payload, err := utils.BodyParser[taskPayloadModel](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = s.tasksService.Patch(r.Context(), uint(taskId), db.Set(
		db.Field("name", payload.Name), // NOTE: this is bad
		db.Field("list_id", payload.ListId),
	)); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	task, err := s.tasksService.Get(r.Context(), uint(taskId))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, task)
}

func (s *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	taskId := utils.UrlNumericParam(w, r, "id")
	if err := s.tasksService.Delete(r.Context(), uint(taskId)); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, nil, "Task deleted")
}
