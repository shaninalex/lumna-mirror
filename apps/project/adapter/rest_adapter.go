// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package adapter

import (
	"github.com/gofiber/fiber/v2"
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
func NewPatchTaskInput(ctx *fiber.Ctx) (*PatchTaskInput, error) {
	projectCode := ctx.Params("projectCode")
	taskCode := ctx.Params("taskCode")

	data, err := web.ParseBody[dto.ChangeTaskStatusDTO](ctx)
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
func NewUpdateTaskInput(ctx *fiber.Ctx) (*UpdateTaskInput, error) {
	projectCode := ctx.Params("projectCode")
	taskCode := ctx.Params("taskCode")
	data, err := web.ParseBody[UpdateTaskData](ctx)
	if err != nil {
		return nil, err
	}

	return &UpdateTaskInput{
		ProjectCode: projectCode,
		TaskCode:    taskCode,
		Data:        data,
	}, nil
}
