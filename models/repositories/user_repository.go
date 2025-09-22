// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/shaninalex/flowreon/internal/utils"
	"gitlab.com/shaninalex/flowreon/models"
)

// UserGetByField get user by field
func UserGetByField(ctx context.Context, db *sql.DB, field string, value any) (*models.User, error) {
	user := &models.User{}
	query := fmt.Sprintf(`
	SELECT id, email, settings, active, state, code, password_hash, created_at, updated_at
	FROM users WHERE %s = ? LIMIT 1
	`, field)
	row := db.QueryRowContext(ctx, query, value)
	err := row.Scan(&user.ID, &user.Email, &user.Settings, &user.Active, &user.State, &user.Code, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}

// UserSave save user
func UserSave(ctx context.Context, db *sql.DB, user *models.User) (*models.User, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}
	user.SetSettings(&models.DefaultUserSettings)
	user.Active = false
	user.State = models.UserStatePending
	user.Code = utils.GenerateEntityCode("user")
	query := `
	INSERT INTO users (email, settings, code, password_hash)
	VALUES (?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, &user.Email, &user.Settings, &user.Code, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserUpdate update user
func UserUpdate(ctx context.Context, db *sql.DB, user *models.User) error {
	query := `
		UPDATE users
		SET 
		    email = ?, 
		    settings = ?, 
		    active = ?, 
		    state = ?, 
		    code = ?,
			updated_at = ?
		WHERE id = ?
	`
	_, err := db.ExecContext(ctx, query,
		user.Email,
		user.Settings,
		user.Active,
		user.State,
		user.Code,
		user.UpdatedAt,
	)
	return err
}
