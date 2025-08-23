package apperrors

import (
	"fmt"
)

type AppError struct {
	ID      string `json:"id"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.ID, e.Message)
}

var (
	Default              = AppError{"APP000", "generic_error", "Something went wrong"}
	InvalidCredentials   = AppError{"AUTH001", "invalid_credentials", "Invalid username or password"}
	TokenExpired         = AppError{"AUTH002", "token_expired", "Authentication token expired"}
	SessionNotFound      = AppError{"AUTH003", "session_not_found", "Session not found"}
	UserNotFound         = AppError{"USER001", "user_not_found", "User not found"}
	UserNotActive        = AppError{"USER002", "user_not_active", "User is not active"}
	UserIdentityNotFound = AppError{"USER003", "identity_not_found", "Identity not found"}
	UserOrgNotAttached   = AppError{"USER004", "org_not_attached", "user does not attach to any organizations"}
	DBConnectionFailed   = AppError{"DB001", "db_connection_failed", "Database connection failed"}
	OrgNotFound          = AppError{"ORG001", "org_not_found", "Organization not found"}
	ProjectNotFound      = AppError{"PRJ001", "prj_not_found", "Project not found"}
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
