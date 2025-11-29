package repositories

import "context"

type Repository[T any] interface {
	NewObject() T
	Get(ctx context.Context, id uint) (*T, error)
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, entry *T) error
	List(ctx context.Context, options map[string]any) []*T
	Update(ctx context.Context, entry *T, columns map[string]any) error
	Count(ctx context.Context, options map[string]any) (int, error)
}
