package utils

import (
	"encoding/json"
	"net/http"
)

// APIResponse standard response structure
type APIResponse[T any] struct {
	Status   bool     `json:"status"`
	Data     T        `json:"data,omitempty"`
	Messages []string `json:"messages,omitempty"`

	// TODO: make app error type
	Errors []error `json:"errors,omitempty"`
}

// NewAPIResponse - new api response.
func NewAPIResponse[T any](data T) *APIResponse[T] {
	return &APIResponse[T]{
		Status: true,
		Data:   data,
	}
}

// ReturnJSON writes JSON response
func ReturnJSON(w http.ResponseWriter, status int, data any, params ...any) {
	resp := NewAPIResponse(data)
	if status >= 400 {
		resp.Status = false
	}

	for _, p := range params {
		switch v := p.(type) {
		case string:
			resp.Messages = append(resp.Messages, v)
		case error:
			resp.Errors = append(resp.Errors, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

// Success is a shorthand for 200 OK
func Success(w http.ResponseWriter, data any, params ...any) {
	ReturnJSON(w, http.StatusOK, data, params...)
}

// Error is a shorthand for error responses
func Error(w http.ResponseWriter, status int, err error) {
	ReturnJSON(w, status, nil, err)
}
