package repositories

import "context"

type Option struct {
	Key   string
	Value any
}

type Repository[T any] interface {
	Get(ctx context.Context, id uint) (*T, error)
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, entry *T) error
	List(ctx context.Context, where string) ([]*T, error)
	Update(ctx context.Context, entry *T, opts ...Option) error
	Count(ctx context.Context, where string) (int, error)
}
