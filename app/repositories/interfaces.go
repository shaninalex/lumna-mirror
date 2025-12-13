package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type Repository[T any] interface {
	Get(ctx context.Context, id uint) (*T, error)
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, entry *T) error
	List(ctx context.Context, opts ...db.Option) ([]*T, error)
	Update(ctx context.Context, entry *T, opts ...db.Option) error
	Count(ctx context.Context, opts ...db.Option) (int, error)
}
