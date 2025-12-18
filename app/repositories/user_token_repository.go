package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type UserTokenRepository struct {
}

func NewUserTokenRepository() *UserTokenRepository {
	return &UserTokenRepository{}
}

var _ Repository[models.UserToken] = (*UserTokenRepository)(nil)

func (s *UserTokenRepository) Get(ctx context.Context, id uint) (*models.UserToken, error) {
	return nil, nil
}

func (s *UserTokenRepository) Delete(ctx context.Context, id uint) error {
	query := `
		DELETE FROM 
			users_tokens
		WHERE 
		    id = ?
	`
	res, err := db.FromContext(ctx).ExecContext(ctx, query, id)
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

func (s *UserTokenRepository) Create(ctx context.Context, entry *models.UserToken) error {
	query := `
		INSERT INTO users_tokens 
		    (user_id, refresh_token, refresh_expires_at)
		VALUES 
		    (?, ?, ?)
	`
	result, err := db.FromContext(ctx).ExecContext(ctx, query,
		&entry.UserId,
		&entry.RefreshToken,
		&entry.RefreshExpiresAt,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	entry.SetId(uint(id))
	return err
}

func (s *UserTokenRepository) List(ctx context.Context, opts db.Expr) ([]*models.UserToken, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT id, user_id, refresh_token, refresh_expires_at, revoked, revoked_at, created_at FROM users_tokens %s`,
		where,
	)

	rows, err := db.FromContext(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := make([]*models.UserToken, 0)
	for rows.Next() {
		token := &models.UserToken{}
		if err = rows.Scan(
			&token.Id,
			&token.UserId,
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

func (s *UserTokenRepository) Update(ctx context.Context, id uint, opts db.SetExpr) error {
	opts.Append(db.Field("updated_at", time.Now()))
	setQuery, args := opts.Build()
	query := fmt.Sprintf(
		`UPDATE user_tokens %s WHERE id = ?`,
		setQuery,
	)

	args = append(args, id)

	res, err := db.FromContext(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}

func (s *UserTokenRepository) Count(ctx context.Context, opts db.Expr) (int, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT count(*) FROM users_tokens %s`,
		where,
	)

	var count int
	row := db.FromContext(ctx).QueryRowContext(ctx, query, args...)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// GetTokenByField loads a token from DB by field.
func (s *UserTokenRepository) GetTokenByField(ctx context.Context, field string, value any) (*models.UserToken, error) {
	token := &models.UserToken{}
	query := fmt.Sprintf(`
		SELECT 
		    id, user_id, refresh_token, refresh_expires_at, revoked, revoked_at, created_at
		FROM users_tokens 
		WHERE %s = ? 
		LIMIT 1
	`, field)
	row := db.FromContext(ctx).QueryRowContext(ctx, query, value)
	err := row.Scan(
		&token.Id,
		&token.UserId,
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

// DeleteTokenByRefreshString removes a specific token for a user by token Id.
func (s *UserTokenRepository) DeleteTokenByRefreshString(ctx context.Context, userID uint, refreshToken string) error {
	query := `
		DELETE FROM 
			users_tokens
		WHERE 
		    user_id = ? AND refresh_token = ?
	`
	res, err := db.FromContext(ctx).ExecContext(ctx, query, userID, refreshToken)
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

// RevokeToken removes a specific token for a user by token Id.
func (s *UserTokenRepository) RevokeToken(ctx context.Context, userId, tokenId uint) error {
	return s.Update(
		ctx,
		tokenId,
		db.Set(
			db.Field("revoked", true),
			db.Field("revoked_at", time.Now()),
		),
	)
}
