package app

import (
	"context"

	"gitlab.com/shaninalex/jajirra/internal/domain"
)

type IAuthApi interface {
	HookRegister(ctx context.Context, data *domain.HooksKratosPayloadDTO) error
	HookVerify(ctx context.Context) error
	HookLogin(ctx context.Context) error
}
