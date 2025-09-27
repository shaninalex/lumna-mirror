// Copyright © 2025 Lumna. All rights reserved.

package handler

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// TaskHandler - task handler.
type TaskHandler struct {
	manager domain.ProjectManager
}

// NewTaskHandler - new task handler.
func NewTaskHandler(manager domain.ProjectManager) *TaskHandler {
	return &TaskHandler{
		manager: manager,
	}
}

// HandleTaskPatchStatus - handle task patch status.
func (s *TaskHandler) HandleTaskPatchStatus(w http.ResponseWriter, r *http.Request) {
	in, err := adapter.NewPatchTaskInput(r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = s.manager.PatchTaskStatus(r.Context(), in.TaskCode, in.Data); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "Task saved")
}

// HandleTaskDetail - handle task detail.
func (s *TaskHandler) HandleTaskDetail(w http.ResponseWriter, r *http.Request) {
	taskCode := r.PathValue("taskCode")
	task, err := s.manager.TaskDetail(r.Context(), taskCode)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewTaskDto(task))
}

// HandleProjectTasksList - handle project tasks list.
func (s *TaskHandler) HandleProjectTasksList(w http.ResponseWriter, r *http.Request) {
	projectCode := r.PathValue("projectCode")

	tasks, err := s.manager.TasksList(r.Context(), projectCode)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewTasksDto(tasks))
}

// HandleTaskUpdate - handle task update.
func (s *TaskHandler) HandleTaskUpdate(w http.ResponseWriter, r *http.Request) {
	data, err := adapter.NewUpdateTaskInput(r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = s.manager.TaskUpdate(r.Context(), data.TaskCode, data.Data); err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, nil, "Task saved")
}

// HandleTaskCreate - create task handler
func (s *TaskHandler) HandleTaskCreate(w http.ResponseWriter, r *http.Request) {
	data, err := web.BodyParser[dto.CreateTaskDto](r)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	task, err := s.manager.TaskCreate(r.Context(), web.GetUserID(r), data)
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	web.Success(w, adapter.NewTaskDto(task), "Task created")
}
