package base

type ContextKey string

const (
	ContextSession = "session"
	ContextUser    = "user"
	ContextUserID  = "user"
	ContextDB      = "postgres_database"
)
