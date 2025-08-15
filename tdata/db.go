// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

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
