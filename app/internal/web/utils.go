package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"gitlab.com/shaninalex/lumna/app/internal/apperrors"
	"gitlab.com/shaninalex/lumna/app/internal/base"
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
func GetUserID(r *http.Request) int64 {
	if id, ok := r.Context().Value(base.ContextUserID).(int64); ok {
		return id
	}
	panic("user was not found in request")
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

func UrlNumericParam(w http.ResponseWriter, r *http.Request, name string) int64 {
	id, err := strconv.ParseInt(r.PathValue(name), 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return 0
	}
	return id
}
