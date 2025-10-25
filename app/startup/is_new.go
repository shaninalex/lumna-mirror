// Copyright © 2025 Lumna. All rights reserved.

package startup

import (
	"database/sql"
)

func IsNew(db *sql.DB) bool {
	row := db.QueryRow(`SELECT count(*) FROM users`)
	var count int
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}

	return count == 0
}
