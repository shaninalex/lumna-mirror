// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"

	"gitlab.com/shaninalex/flowreon/database"
)

func clearDatabase(ctx context.Context) {
	_db := database.GetDB(ctx)
	if err := _db.Exec(`
		TRUNCATE TABLE 
			epics,
		    organizations,
		    projects,
		    sprints,
		    task_statuses,
		    tasks,
		    users
		RESTART IDENTITY CASCADE
	`).Error; err != nil {
		panic(err)
	}
}
