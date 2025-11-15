package handler

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/apps/project/adapter"
	"gitlab.com/shaninalex/lumna/app/domain"
	"gitlab.com/shaninalex/lumna/app/pkg/web"
)

// ProjectTaskHandler - task handler.
type ProjectTaskHandler struct {
	taskManager domain.TaskManager
}

// NewProjectTaskHandler - new task handler.
func NewProjectTaskHandler() *ProjectTaskHandler {
	return &ProjectTaskHandler{
		taskManager: domain.NewTaskService(),
	}
}

// List - retrieve tasks for a project
func (s *ProjectTaskHandler) List(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	tasks, err := s.taskManager.TasksList(r.Context(), projectID)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.ToTaskListDto(tasks))
}

// Create - create a new task in a project
func (s *ProjectTaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	projectID := web.UrlNumericParam(w, r, "id")
	input, err := web.BodyParser[adapter.ProjectInput](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	userID := web.GetUserID(r)

	task := &domain.Task{
		Title:     input.Title,
		ProjectID: projectID,
		StatusID:  input.StatusId,
		UserID:    userID,
	}
	task, err = s.taskManager.TaskCreate(ctx, task)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.ToTaskDto(task))
}
