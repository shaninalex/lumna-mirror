// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

package database

import (
	"context"
	"errors"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IObject interface {
	GetID() uint
	SetID(id uint)
}

type Repository[T IObject] struct {
	DB *gorm.DB
}

// NewRepository creates a new GORM repository instance
func NewRepository[T IObject](db *gorm.DB) *Repository[T] {
	return &Repository[T]{DB: db}
}

// GetByID retrieves a record by ID
func (r *Repository[T]) GetByID(ctx context.Context, id uuid.UUID) (T, error) {
	obj := r.newObject()
	err := GetDB(ctx).WithContext(ctx).
		Where("id = ?", id).
		First(obj).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.returnNil(), err
		}
		return r.returnNil(), err
	}

	return obj, nil
}

// Create inserts a new record
func (r *Repository[T]) Create(ctx context.Context, obj T) (T, error) {
	if err := GetDB(ctx).WithContext(ctx).Create(obj).Error; err != nil {
		return r.returnNil(), err
	}

	return obj, nil
}

// Update modifies an existing record
func (r *Repository[T]) Update(ctx context.Context, obj T) (T, error) {
	if err := GetDB(ctx).WithContext(ctx).
		Model(obj).
		Where("id = ?", obj.GetID()).
		Updates(obj).Error; err != nil {
		return r.returnNil(), err
	}

	return obj, nil
}

// List retrieves multiple records with optional query customization
func (r *Repository[T]) List(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	var results []T
	dbQuery := GetDB(ctx).WithContext(ctx).Model(new(T))

	for _, opt := range opts {
		dbQuery = opt(dbQuery)
	}

	if err := dbQuery.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// GetBy retrieves a single record with custom query options
func (r *Repository[T]) GetBy(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) (T, error) {
	obj := r.newObject()
	dbQuery := GetDB(ctx).WithContext(ctx).Model(obj)

	for _, opt := range opts {
		dbQuery = opt(dbQuery)
	}

	if err := dbQuery.First(obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.returnNil(), err
		}
		return r.returnNil(), err
	}

	return obj, nil
}

// DeleteByID removes a record by ID
func (r *Repository[T]) DeleteByID(ctx context.Context, id uuid.UUID) error {
	obj := r.newObject()
	if err := GetDB(ctx).WithContext(ctx).
		Where("id = ?", id).
		Delete(obj).Error; err != nil {
		return err
	}
	return nil
}

// newObject creates a new zero-value instance of T
func (r *Repository[T]) newObject() T {
	var zero T
	typ := reflect.TypeOf(zero)
	if typ.Kind() == reflect.Pointer {
		return reflect.New(typ.Elem()).Interface().(T)
	}
	return zero
}

// returnNil returns a "nil" value for T
func (r *Repository[T]) returnNil() T {
	var nilVar T
	if any(nilVar) == nil {
		return nilVar
	}
	return nilVar
}
