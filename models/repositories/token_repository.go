// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"gitlab.com/shaninalex/flowreon/models"
)

// GetTokenByField loads a token from DB by field.
func GetTokenByField(ctx context.Context, db *sql.DB, field string, value any) (*models.UserToken, error) {
	token := &models.UserToken{}
	query := fmt.Sprintf(`
		SELECT 
		    id, jti, user_id, device, refresh_token, refresh_expires_at, revoked, revoked_at, created_at
		FROM users_tokens 
		WHERE %s = ? 
		LIMIT 1
	`, field)
	row := db.QueryRowContext(ctx, query, value)
	err := row.Scan(
		&token.ID,
		&token.Jti,
		&token.UserID,
		&token.Device,
		&token.RefreshToken,
		&token.RefreshExpiresAt,
		&token.Revoked,
		&token.RevokedAt,
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
	query := `
		INSERT INTO users_tokens 
		    (user_id, device, jti, refresh_token, refresh_expires_at)
		VALUES 
		    (?, ?, ?, ?, ?)
	`
	result, err := db.ExecContext(ctx, query,
		&token.UserID,
		&token.Device,
		&token.Jti,
		&token.RefreshToken,
		&token.RefreshExpiresAt,
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
		SELECT 
		    id, jti, user_id, device, refresh_token, refresh_expires_at, revoked, revoked_at, created_at
		FROM users_tokens
		WHERE user_id = ?
	`
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
		token := &models.UserToken{}
		if err = rows.Scan(
			&token.ID,
			&token.Jti,
			&token.UserID,
			&token.Device,
			&token.RefreshToken,
			&token.RefreshExpiresAt,
			&token.Revoked,
			&token.RevokedAt,
			&token.CreatedAt,
		); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

// DeleteToken removes a specific token for a user by token ID.
func DeleteToken(ctx context.Context, db *sql.DB, userID uint, jti string) error {
	query := `
		DELETE FROM 
			users_tokens
		WHERE 
		    user_id = ? AND 
		    jti = ?
	`
	res, err := db.ExecContext(ctx, query, userID, jti)
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
