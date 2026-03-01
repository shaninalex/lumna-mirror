package auth

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal"
)

func GetIdentityId(ctx context.Context) uint {
	id, ok := ctx.Value(internal.ContextIdentity).(uint)
	if !ok {
		panic("identity not found in context")
	}
	return id
}
