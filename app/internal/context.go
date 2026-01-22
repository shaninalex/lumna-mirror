package internal

// ContextKey - context key.
type ContextKey string

const (
	// ContextUserID - user id
	ContextUserID ContextKey = "user_id"

	// ContextDB - database
	ContextDB ContextKey = "database"

	// ContextAppName - used to map request to a service
	ContextAppName ContextKey = "app_name"

	ContextConfig ContextKey = "config"

	// ContextIdentity - user id
	ContextIdentity ContextKey = "identity"
)
