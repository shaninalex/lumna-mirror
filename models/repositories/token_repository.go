// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/shaninalex/flowreon/models"
)

// GetTokenByField loads a token from DB by field.
func GetTokenByField(ctx context.Context, db *sql.DB, field string, value any) (*models.UserToken, error) {
	token := &models.UserToken{}
	query := fmt.Sprintf(`
	SELECT id, user_id, claims, device, expires_at, created_at
	FROM users_tokens WHERE %s = ? LIMIT 1`, field)
	row := db.QueryRowContext(ctx, query, value)
	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.Claims,
		&token.Device,
		&token.ExpiresAt,
		&token.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return token, nil
}

// SaveToken inserts or updates a session in the DB.
func SaveToken(ctx context.Context, db *sql.DB, token *models.UserToken) error {
	if token.ExpiresAt.IsZero() {
		token.ExpiresAt = time.Now().Add(7 * 24 * time.Hour) // default 7 days
	}
	query := `
	INSERT INTO users_tokens (user_id, claims, device, expires_at)
	VALUES (?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, query,
		&token.UserID,
		&token.Claims,
		&token.Device,
		&token.ExpiresAt,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	token.SetID(uint(id))
	return err
}

// GetTokens retrieves all tokens for a given user.
func GetTokens(ctx context.Context, db *sql.DB, userID uint) ([]*models.UserToken, error) {
	query := `
	SELECT id, user_id, claims, device, expires_at, created_at
	FROM users_tokens
	WHERE user_id = ?`
	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			panic(err)
		}
	}()

	tokens := make([]*models.UserToken, 0)
	for rows.Next() {
		t := &models.UserToken{}
		if err = rows.Scan(&t.ID, &t.UserID, &t.Claims, &t.Device, &t.ExpiresAt, &t.CreatedAt); err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}
	return tokens, nil
}

// DeleteToken removes a specific token for a user by token ID.
func DeleteToken(ctx context.Context, db *sql.DB, userID, tokenID uint) error {
	query := `
	DELETE FROM users_tokens
	WHERE user_id = ? AND id = ?`
	res, err := db.ExecContext(ctx, query, userID, tokenID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
