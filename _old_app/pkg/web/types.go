package web

import (
	"gitlab.com/shaninalex/lumna/_old_app/pkg/apperrors"
)

// APIResponse standard response structure
type APIResponse[T any] struct {
	Status   bool                 `json:"status"`
	Data     T                    `json:"data,omitempty"`
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
