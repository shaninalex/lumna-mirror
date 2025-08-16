package apperrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppError struct {
	ID         string `json:"id"`
	Key        string `json:"key"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.ID, e.Message)
}

var (
	InvalidCredentials = AppError{"AUTH001", "invalid_credentials", "Invalid username or password", http.StatusUnauthorized}
	TokenExpired       = AppError{"AUTH002", "token_expired", "Authentication token expired", http.StatusUnauthorized}
	UserNotFound       = AppError{"USER001", "user_not_found", "User not found", http.StatusNotFound}
	UserNotActive      = AppError{"USER002", "user_not_active", "User is not active", http.StatusUnauthorized}
	DBConnectionFailed = AppError{"DB001", "db_connection_failed", "Database connection failed", http.StatusInternalServerError}
)

func WriteError(w http.ResponseWriter, err AppError) {
	// TODO: event send to prometheus
	resp := map[string]any{
		"status":   false,
		"data":     err,
		"messages": make([]string, 0),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatus)
	_ = json.NewEncoder(w).Encode(resp)
}
