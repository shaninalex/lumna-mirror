package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse standard response structure
type APIResponse[T any] struct {
	Status   bool     `json:"status"`
	Data     T        `json:"data,omitempty"`
	Messages []string `json:"messages,omitempty"`

	Errors []string `json:"errors,omitempty"`
}

// NewAPIResponse - new api response.
func NewAPIResponse[T any](data T) *APIResponse[T] {
	return &APIResponse[T]{
		Status: true,
		Data:   data,
	}
}

// Success is a shorthand for 200 OK
func Success(c *gin.Context, data any, params ...any) {
	ReturnJSON(c, http.StatusOK, data, params...)
}

// Error is a shorthand for error responses
func Error(c *gin.Context, status int, err error) {
	ReturnJSON(c, status, nil, err)
}

// ReturnJSON writes JSON response
func ReturnJSON(c *gin.Context, status int, data any, params ...any) {
	resp := NewAPIResponse(data)
	if status >= 400 {
		resp.Status = false
	}

	for _, p := range params {
		switch v := p.(type) {
		case string:
			resp.Messages = append(resp.Messages, v)
		case error:
			resp.Errors = append(resp.Errors, v.Error())
		}
	}

	c.JSON(status, resp)
}
