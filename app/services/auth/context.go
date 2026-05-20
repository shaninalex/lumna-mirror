package auth

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/pkg"
)

func GetIdentityId(ctx context.Context) uint {
	id, ok := ctx.Value(pkg.ContextIdentity).(uint)
	if !ok {
		panic("identity not found in context")
	}
	return id
}
