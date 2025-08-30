// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package base

// ContextKey - context key.
type ContextKey string

const (
	ContextSession = "session"
	ContextUser    = "user"
	ContextUserID  = "user_id"
	ContextOrgID   = "org_id"
	ContextDB      = "postgres_database"
)
