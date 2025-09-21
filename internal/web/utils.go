// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/base"
)

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
		case apperrors.AppError:
			resp.Errors = append(resp.Errors, v)
		case error:
			var appErr apperrors.AppError
			if errors.As(v, &appErr) {
				resp.Errors = append(resp.Errors, appErr)
			} else {
				resp.Messages = append(resp.Messages, v.Error())
			}
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

// GetUserID retrieves the user ID from context
func GetUserID(r *http.Request) uuid.UUID {
	if id, ok := r.Context().Value(base.ContextUserID).(string); ok {
		return uuid.MustParse(id)
	}
	return uuid.Nil
}

// GetOrganizationID retrieves the organization ID from context
func GetOrganizationID(r *http.Request) uuid.UUID {
	if id, ok := r.Context().Value(base.ContextOrgID).(uuid.UUID); ok {
		return id
	}
	return uuid.Nil
}

// GetAppName retrieves the app name from context
func GetAppName(r *http.Request) *string {
	if user, ok := r.Context().Value(base.ContextAppName).(*string); ok {
		return user
	}
	return nil
}

// BodyParser parse request POST body into generic type
func BodyParser[T any](r *http.Request) (*T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
