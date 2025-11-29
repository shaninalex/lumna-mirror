package apperrors

import (
	"fmt"
)

// AppError - app error.
type AppError struct {
	Id      string `json:"id"`
	Key     string `json:"key"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Error implements error interface
func (e AppError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.Id, e.Message)
}

var (
	// Default default application error
	Default = AppError{"APP000", "generic_error", "Something went wrong", nil}

	// InvalidCredentials - Invalid Credentials error
	InvalidCredentials = AppError{"AUTH001", "invalid_credentials", "Invalid username or password", nil}

	// TokenExpired - Token Expired error
	TokenExpired = AppError{"AUTH002", "token_expired", "Authentication token expired", nil}

	// SessionNotFound - Session Not Found error
	SessionNotFound = AppError{"AUTH003", "session_not_found", "Session not found", nil}

	// UserNotFound - User Not Found error
	UserNotFound = AppError{"USER001", "user_not_found", "User not found", nil}

	// UserNotActive - User Not Active error
	UserNotActive = AppError{"USER002", "user_not_active", "User is not active", nil}

	// UserIdentityNotFound - User Identity Not Found error
	UserIdentityNotFound = AppError{"USER003", "identity_not_found", "Identity not found", nil}

	// UserOrgNotAttached - User Org Not Attached error
	UserOrgNotAttached = AppError{"USER004", "org_not_attached", "user does not attach to any organizations", nil}

	// UserUnableToCreate - User Unable To Create error
	UserUnableToCreate = AppError{"USER005", "unable_to_create", "unable to create user", nil}

	// DBConnectionFailed - DB Connection Failed error
	DBConnectionFailed = AppError{"DB001", "db_connection_failed", "Database connection failed", nil}

	// OrgNotFound - Org Not Found error
	OrgNotFound = AppError{"ORG001", "org_not_found", "Organization not found", nil}

	// ProjectNotFound - Project Not Found error
	ProjectNotFound = AppError{"PRJ001", "prj_not_found", "Project not found", nil}

	// TaskNotFound - Task Not Found error
	TaskNotFound = AppError{"TAS001", "task_not_found", "Task not found", nil}

	// TaskUnableToPatch - Task Unable To Patch error
	TaskUnableToPatch = AppError{"TAS002", "unable_to_patch", "Unable to patch task", nil}
)

// AllErrors set of all application errors
var AllErrors = []AppError{
	Default,
	InvalidCredentials,
	TokenExpired,
	SessionNotFound,
	UserNotFound,
	UserNotActive,
	UserIdentityNotFound,
	UserOrgNotAttached,
	UserUnableToCreate,
	DBConnectionFailed,
	OrgNotFound,
	ProjectNotFound,
	TaskNotFound,
	TaskUnableToPatch,
}
