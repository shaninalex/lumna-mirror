package base

type ContextKey string

const (
	ContextSession = "session"
	ContextUser    = "user"
	ContextUserID  = "user_id"
	ContextOrgID   = "org_id"
	ContextDB      = "postgres_database"
)
