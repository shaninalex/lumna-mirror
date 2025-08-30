// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
)

// APIResponse - api response.
type APIResponse[T any] struct {
	Status   bool                 `json:"status"`
	Data     T                    `json:"data"`
	Messages []string             `json:"messages,omitempty"`
	Errors   []apperrors.AppError `json:"errors,omitempty"`
}

// NewAPIResponse - new api response.
func NewAPIResponse[T any](data T) *APIResponse[T] {
	return &APIResponse[T]{
		Status: true,
		Data:   data,
	}
}
