package app

import "context"

type IAuthApi interface {
	Register(ctx context.Context)
	Verify(ctx context.Context)
	Login(ctx context.Context)
	Restore(ctx context.Context)
	Session(ctx context.Context)
}
