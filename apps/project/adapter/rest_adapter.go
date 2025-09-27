// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// PatchTaskInput - patch task input.
type PatchTaskInput struct {
	ProjectCode string
	TaskCode    string
	Data        *dto.ChangeTaskStatusDTO
}

// NewPatchTaskInput - new patch task input.
func NewPatchTaskInput(r *http.Request) (*PatchTaskInput, error) {
	projectCode := r.PathValue("projectCode")
	taskCode := r.PathValue("taskCode")

	data, err := web.BodyParser[dto.ChangeTaskStatusDTO](r)
	if err != nil {
		return nil, err
	}

	return &PatchTaskInput{
		ProjectCode: projectCode,
		TaskCode:    taskCode,
		Data:        data,
	}, nil
}

// UpdateTaskData - updates the task data.
type UpdateTaskData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateTaskInput - updates the task input.
type UpdateTaskInput struct {
	ProjectCode string
	TaskCode    string
	Data        *UpdateTaskData
}

// NewUpdateTaskInput - new update task input.
func NewUpdateTaskInput(r *http.Request) (*UpdateTaskInput, error) {
	projectCode := r.PathValue("projectCode")
	taskCode := r.PathValue("taskCode")
	data, err := web.BodyParser[UpdateTaskData](r)
	if err != nil {
		return nil, err
	}

	return &UpdateTaskInput{
		ProjectCode: projectCode,
		TaskCode:    taskCode,
		Data:        data,
	}, nil
}
