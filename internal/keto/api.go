package keto

import (
	"context"

	"github.com/google/uuid"
)

type IKeto interface {
	GetPermissionsTree(ctx context.Context, userID uuid.UUID) any
}
