// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/models"
)

// GetKratosRedirectURL returns redirect URL using Kratos base URL from config
func GetKratosRedirectURL(c base.IConfig, path string) string {
	return fmt.Sprintf("%s%s", c.String("kratos.url_browser"), path)
}

// NewAPIResponse initializes a response
func NewAPIResponse[T any](data any) *APIResponse[T] {
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
	if id, ok := r.Context().Value(base.ContextUserID).(uuid.UUID); ok {
		return id
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

// GetUser retrieves the user object from context
func GetUser(r *http.Request) *models.User {
	if user, ok := r.Context().Value(base.ContextUser).(*models.User); ok {
		return user
	}
	return nil
}

// GetSession retrieves the session object from context
func GetSession(r *http.Request) *ory.Session {
	if session, ok := r.Context().Value(base.ContextSession).(*ory.Session); ok {
		return session
	}
	return nil
}

// ParamUUID parses a URL param as UUID
func ParamUUID(r *http.Request, name string, params map[string]string) (uuid.UUID, error) {
	val, ok := params[name]
	if !ok {
		return uuid.Nil, fmt.Errorf("%s is required", name)
	}
	id, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s is not a valid UUID", name)
	}
	return id, nil
}

// ParamString retrieves a string param
func ParamString(r *http.Request, name string, params map[string]string) (string, error) {
	val, ok := params[name]
	if !ok || val == "" {
		return "", fmt.Errorf("%s is required", name)
	}
	return val, nil
}

// BodyParser parse request POST body into generic type
func BodyParser[T any](r *http.Request) (*T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
