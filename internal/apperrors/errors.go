// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package apperrors

import (
	"fmt"
)

// AppError - app error.
type AppError struct {
	ID      string `json:"id"`
	Key     string `json:"key"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Error implements error interface
func (e AppError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.ID, e.Message)
}

var (
	Default              = AppError{"APP000", "generic_error", "Something went wrong", nil}
	InvalidCredentials   = AppError{"AUTH001", "invalid_credentials", "Invalid username or password", nil}
	TokenExpired         = AppError{"AUTH002", "token_expired", "Authentication token expired", nil}
	SessionNotFound      = AppError{"AUTH003", "session_not_found", "Session not found", nil}
	UserNotFound         = AppError{"USER001", "user_not_found", "User not found", nil}
	UserNotActive        = AppError{"USER002", "user_not_active", "User is not active", nil}
	UserIdentityNotFound = AppError{"USER003", "identity_not_found", "Identity not found", nil}
	UserOrgNotAttached   = AppError{"USER004", "org_not_attached", "user does not attach to any organizations", nil}
	UserUnableToCreate   = AppError{"USER005", "unable_to_create", "unable to create user", nil}
	DBConnectionFailed   = AppError{"DB001", "db_connection_failed", "Database connection failed", nil}
	OrgNotFound          = AppError{"ORG001", "org_not_found", "Organization not found", nil}
	ProjectNotFound      = AppError{"PRJ001", "prj_not_found", "Project not found", nil}
	TaskNotFound         = AppError{"TAS001", "task_not_found", "Task not found", nil}
	TaskUnableToPatch    = AppError{"TAS002", "unable_to_patch", "Unable to patch task", nil}
)

var AllErrors = []AppError{
	Default,
	InvalidCredentials,
	TokenExpired,
	SessionNotFound,
	UserNotFound,
	UserNotActive,
	UserIdentityNotFound,
	UserOrgNotAttached,
	DBConnectionFailed,
	OrgNotFound,
	ProjectNotFound,
}
