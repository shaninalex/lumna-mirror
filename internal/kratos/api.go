package kratos

import (
	"context"

	"github.com/google/uuid"
	kratos "github.com/ory/kratos-client-go"
)

type IKratos interface {
	Get(ctx context.Context, id uuid.UUID) kratos.Identity
}
