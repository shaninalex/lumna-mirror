// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package base

// ContextKey - context key.
type ContextKey string

const (
	// ContextSession ory session context key
	ContextSession = "session"

	// ContextUser - user object
	ContextUser = "user"

	// ContextUserID - user id
	ContextUserID = "user_id"

	// ContextOrgID - user organization id
	ContextOrgID = "org_id"

	// ContextDB - gorm db
	ContextDB = "postgres_database"

	// ContextCookie - authentication cookie
	ContextCookie = "auth_cookie"

	// ContextAppName - used to map request to a service
	ContextAppName = "app_name"
)
