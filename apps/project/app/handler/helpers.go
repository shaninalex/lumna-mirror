// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

type PatchTaskInput struct {
	ProjectCode string
	TaskCode    string
	Data        *domain.ChangeTaskStatusDTO
}

func NewPatchTaskInput(ctx *fiber.Ctx) (*PatchTaskInput, error) {
	projectCode := ctx.Params("projectCode")
	taskCode := ctx.Params("taskCode")

	data, err := web.ParseBody[domain.ChangeTaskStatusDTO](ctx)
	if err != nil {
		return nil, err
	}

	return &PatchTaskInput{
		ProjectCode: projectCode,
		TaskCode:    taskCode,
		Data:        data,
	}, nil
}

type UpdateTaskInput struct {
	ProjectCode string
	TaskCode    string
	Data        *domain.UpdateTaskData
}

func NewUpdateTaskInput(ctx *fiber.Ctx) (*UpdateTaskInput, error) {
	projectCode := ctx.Params("projectCode")
	taskCode := ctx.Params("taskCode")
	data, err := web.ParseBody[domain.UpdateTaskData](ctx)
	if err != nil {
		return nil, err
	}

	return &UpdateTaskInput{
		ProjectCode: projectCode,
		TaskCode:    taskCode,
		Data:        data,
	}, nil
}
