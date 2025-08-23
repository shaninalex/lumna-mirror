package apperrors

import (
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
	Default            = AppError{"APP000", "generic_error", "Something went wrong", http.StatusBadRequest}
	InvalidCredentials = AppError{"AUTH001", "invalid_credentials", "Invalid username or password", http.StatusUnauthorized}
	TokenExpired       = AppError{"AUTH002", "token_expired", "Authentication token expired", http.StatusUnauthorized}
	UserNotFound       = AppError{"USER001", "user_not_found", "User not found", http.StatusNotFound}
	UserNotActive      = AppError{"USER002", "user_not_active", "User is not active", http.StatusUnauthorized}
	DBConnectionFailed = AppError{"DB001", "db_connection_failed", "Database connection failed", http.StatusInternalServerError}
	OrgNotFound        = AppError{"ORG001", "org_not_found", "Organization not found", http.StatusNotFound}
	ProjectNotFound    = AppError{"PRJ001", "prj_not_found", "Project not found", http.StatusNotFound}
)

var AllErrors = []AppError{
	Default,
	InvalidCredentials,
	TokenExpired,
	UserNotFound,
	UserNotActive,
	DBConnectionFailed,
	OrgNotFound,
	ProjectNotFound,
}
