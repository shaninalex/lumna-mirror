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
	List(ctx context.Context, options map[string]any) []*T
	Update(ctx context.Context, entry *T, opts ...Option) error
	Count(ctx context.Context, options map[string]any) (int, error)
}
