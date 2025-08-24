package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/internal/domain"
	"gitlab.com/shaninalex/jajirra/internal/web"
)

type PatchTaskInput struct {
	ProjectKey string
	TaskID     uuid.UUID
	Data       *domain.ChangeTaskStatusDTO
}

func NewPatchTaskInput(ctx *fiber.Ctx) (*PatchTaskInput, error) {
	projectKey, err := web.ParamString(ctx, "projectKey")
	if err != nil {
		return nil, err
	}
	taskID, err := web.ParamUUID(ctx, "taskID")
	if err != nil {
		return nil, err
	}

	data, err := web.ParseBody[domain.ChangeTaskStatusDTO](ctx)
	if err != nil {
		return nil, err
	}

	return &PatchTaskInput{
		ProjectKey: projectKey,
		TaskID:     taskID,
		Data:       data,
	}, nil
}
