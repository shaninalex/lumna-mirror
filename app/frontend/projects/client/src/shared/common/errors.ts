export interface AppError {
	id: string
	key: string
	message: string
    data?: any
}

export const ERRORS: Record<string, AppError> = {
	"db_connection_failed": {
		"id": "DB001",
		"key": "db_connection_failed",
		"message": "Database connection failed"
	},
	"generic_error": {
		"id": "APP000",
		"key": "generic_error",
		"message": "Something went wrong"
	},
	"identity_not_found": {
		"id": "USER003",
		"key": "identity_not_found",
		"message": "Identity not found"
	},
	"invalid_credentials": {
		"id": "AUTH001",
		"key": "invalid_credentials",
		"message": "Invalid username or password"
	},
	"org_not_attached": {
		"id": "USER004",
		"key": "org_not_attached",
		"message": "user does not attach to any organizations"
	},
	"org_not_found": {
		"id": "ORG001",
		"key": "org_not_found",
		"message": "Organization not found"
	},
	"prj_not_found": {
		"id": "PRJ001",
		"key": "prj_not_found",
		"message": "Project not found"
	},
	"session_not_found": {
		"id": "AUTH003",
		"key": "session_not_found",
		"message": "Session not found"
	},
	"task_not_found": {
		"id": "TAS001",
		"key": "task_not_found",
		"message": "Task not found"
	},
	"token_expired": {
		"id": "AUTH002",
		"key": "token_expired",
		"message": "Authentication token expired"
	},
	"unable_to_create": {
		"id": "USER005",
		"key": "unable_to_create",
		"message": "unable to create user"
	},
	"unable_to_patch": {
		"id": "TAS002",
		"key": "unable_to_patch",
		"message": "Unable to patch task"
	},
	"user_not_active": {
		"id": "USER002",
		"key": "user_not_active",
		"message": "User is not active"
	},
	"user_not_found": {
		"id": "USER001",
		"key": "user_not_found",
		"message": "User not found"
	}
}
