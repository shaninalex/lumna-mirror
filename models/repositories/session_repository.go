// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package repositories

import (
	"context"
	"database/sql"
	"time"

	"gitlab.com/shaninalex/flowreon/models"
)

// GetSessionByID loads a session from DB by session ID.
func GetSessionByID(ctx context.Context, db *sql.DB, id string) (*models.SessionModel, error) {
	sm := &models.SessionModel{}
	query := `
	SELECT id, user_id, data, expires_at, created_at
	FROM user_sessions
	WHERE id = ?
	LIMIT 1
	`

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&sm.ID,
		&sm.UserID,
		&sm.Data,
		&sm.ExpiresAt,
		&sm.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return sm, nil
}

// SaveSession inserts or updates a session in the DB.
func SaveSession(ctx context.Context, db *sql.DB, sm *models.SessionModel) error {
	if sm.CreatedAt.IsZero() {
		sm.CreatedAt = time.Now()
	}
	if sm.ExpiresAt.IsZero() {
		sm.ExpiresAt = time.Now().Add(7 * 24 * time.Hour) // default 7 days
	}

	query := `
	INSERT INTO user_sessions (id, user_id, data, expires_at, created_at)
	VALUES (?, ?, ?, ?, ?)
	ON CONFLICT(id) DO UPDATE SET
		data = excluded.data,
		expires_at = excluded.expires_at
	`

	_, err := db.ExecContext(ctx, query,
		sm.ID,
		sm.UserID,
		sm.Data,
		sm.ExpiresAt,
		sm.CreatedAt,
	)
	return err
}
