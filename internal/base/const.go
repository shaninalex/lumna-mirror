package base

type ContextKey string

const (
	ContextSession = "session"
	ContextUser    = "user"
	ContextDB      = "postgres_database"
)
