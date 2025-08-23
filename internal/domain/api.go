package domain

import (
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
)

type ApiResponse struct {
	Status   bool                 `json:"status"`
	Data     any                  `json:"data"`
	Messages []string             `json:"messages,omitempty"`
	Errors   []apperrors.AppError `json:"errors,omitempty"`
}

func NewApiResponse(data any) *ApiResponse {
	return &ApiResponse{
		Status: true,
		Data:   data,
	}
}
