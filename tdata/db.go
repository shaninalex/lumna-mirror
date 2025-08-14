// Copyright © 2025 Soundstream https://soundstream.shaninalex.com. All rights reserved.

package tdata

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTestDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:?cache=shared"))
	if err != nil {
		panic(err)
	}

	return db
}
