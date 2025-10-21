// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

// GetTokenByField loads a token from DB by field.
func GetTokenByField(ctx context.Context, db *sql.DB, field string, value any) (*UserToken, error) {
	token := &UserToken{}
	query := fmt.Sprintf(`
		SELECT 
		    id, user_id, device, refresh_token, refresh_expires_at, revoked, revoked_at, created_at
		FROM users_tokens 
		WHERE %s = ? 
		LIMIT 1
	`, field)
	row := db.QueryRowContext(ctx, query, value)
	err := row.Scan(
		&token.ID,
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
func SaveToken(ctx context.Context, db *sql.DB, token *UserToken) error {
	query := `
		INSERT INTO users_tokens 
		    (user_id, device, refresh_token, refresh_expires_at)
		VALUES 
		    (?, ?, ?, ?)
	`
	result, err := db.ExecContext(ctx, query,
		&token.UserID,
		&token.Device,
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
func GetTokens(ctx context.Context, db *sql.DB, userID uint) ([]*UserToken, error) {
	query := `
		SELECT 
		    id, user_id, device, refresh_token, refresh_expires_at, revoked, revoked_at, created_at
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

	tokens := make([]*UserToken, 0)
	for rows.Next() {
		token := &UserToken{}
		if err = rows.Scan(
			&token.ID,
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
func DeleteToken(ctx context.Context, db *sql.DB, userID, id uint) error {
	query := `
		DELETE FROM 
			users_tokens
		WHERE 
		    user_id = ? AND 
		    id = ?
	`
	res, err := db.ExecContext(ctx, query, userID, id)
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

// DeleteTokenByRefreshString removes a specific token for a user by token ID.
func DeleteTokenByRefreshString(ctx context.Context, db *sql.DB, userID uint, refreshToken string) error {
	query := `
		DELETE FROM 
			users_tokens
		WHERE 
		    user_id = ? AND refresh_token = ?
	`
	res, err := db.ExecContext(ctx, query, userID, refreshToken)
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

// RevokeToken removes a specific token for a user by token ID.
func RevokeToken(ctx context.Context, db *sql.DB, userID, tokenID uint) error {
	query := `
		UPDATE users_tokens
		SET
		    revoked = true,
		    revoked_at = NOW()
		WHERE 
		    user_id = ? AND 
		    refresh_token = ?
	`
	result, err := db.ExecContext(ctx, query, userID, tokenID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
