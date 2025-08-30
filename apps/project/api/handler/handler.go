// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
)

type ProjectHandler struct {
	projectApi *domain.ProjectManagement
}

func NewProjectHandler() *ProjectHandler {
	h := &ProjectHandler{
		projectApi: domain.NewProjectManagement(),
	}
	return h
}

type TaskFilter struct {
	Project  string `query:"project,required"`
	TaskCode string `query:"taskCode"`
}

func (s *ProjectHandler) getFilterParams(ctx *fiber.Ctx) (*TaskFilter, error) {
	q := new(TaskFilter)
	err := ctx.QueryParser(q)
	return q, err
}
