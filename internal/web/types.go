// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
)

// ApiResponse - api response.
type ApiResponse[T any] struct {
	Status   bool                 `json:"status"`
	Data     T                    `json:"data"`
	Messages []string             `json:"messages,omitempty"`
	Errors   []apperrors.AppError `json:"errors,omitempty"`
}

// NewApiResponse - new api response.
func NewApiResponse[T any](data T) *ApiResponse[T] {
	return &ApiResponse[T]{
		Status: true,
		Data:   data,
	}
}
