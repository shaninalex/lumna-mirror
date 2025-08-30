// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
)

// ApiResponse - api response.
type ApiResponse struct {
	Status   bool                 `json:"status"`
	Data     any                  `json:"data"`
	Messages []string             `json:"messages,omitempty"`
	Errors   []apperrors.AppError `json:"errors,omitempty"`
}

// NewApiResponse - new api response.
func NewApiResponse(data any) *ApiResponse {
	return &ApiResponse{
		Status: true,
		Data:   data,
	}
}
