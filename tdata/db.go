// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"

	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/models"
)

func resetDatabase(ctx context.Context) {
	db := database.GetDB(ctx)
	err := db.Migrator().DropTable(
		&models.User{},
		&models.Task{},
		&models.TaskStatus{},
		&models.Epic{},
		&models.Sprint{},
		&models.Organization{},
		&models.Project{},
	)
	if err != nil {
		panic(err)
	}

	database.ApplyMigrations(db)
}
