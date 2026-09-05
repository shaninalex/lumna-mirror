package auth

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/pkg"
)

func GetIdentityId(ctx context.Context) int {
	id, ok := ctx.Value(pkg.ContextIdentity).(int)
	if !ok {
		panic("identity not found in context")
	}
	return id
}
