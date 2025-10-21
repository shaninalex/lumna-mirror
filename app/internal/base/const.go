// Copyright © 2025 Lumna. All rights reserved.

package base

// ContextKey - context key.
type ContextKey string

const (
	// ContextUserID - user id
	ContextUserID = "user_id"

	// ContextDB - database
	ContextDB = "database"

	// ContextAppName - used to map request to a service
	ContextAppName = "app_name"
)
