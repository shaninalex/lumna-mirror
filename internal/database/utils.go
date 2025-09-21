// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package database

import (
	"fmt"

	"gitlab.com/shaninalex/flowreon/internal/base"
)

// BuildDSN - builds the dsn.
func BuildDSN(c base.IConfig) string {
	host := c.String("db.POSTGRES_HOST")
	user := c.String("db.POSTGRES_USER")
	pass := c.String("db.POSTGRES_PASSWORD")
	name := c.String("db.POSTGRES_DB")
	port := c.Int("db.POSTGRES_PORT")

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, pass, host, port, name)
}
