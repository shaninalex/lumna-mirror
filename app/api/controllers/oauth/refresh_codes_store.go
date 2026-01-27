package oauth

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

// type RefreshToken struct {
// 	Token     string
// 	ClientID  string
// 	UserID    string
// 	Scopes    []string
// 	ExpiresAt time.Time
// 	Revoked   bool
// }

type RefreshTokenStore interface {
	Save(ctx context.Context, token *models.RefreshToken) error
	FindByHash(ctx context.Context, hash string) (*models.RefreshToken, error)
	Revoke(ctx context.Context, id uuid.UUID) error
	RevokeAll(ctx context.Context, userID uuid.UUID, clientID string) error
}

func NewPersistentRefreshTokenStore() *PersistentRefreshTokenStore {
	return &PersistentRefreshTokenStore{}
}

var _ RefreshTokenStore = (*PersistentRefreshTokenStore)(nil)

type PersistentRefreshTokenStore struct {
}

// FindByHash implements [RefreshTokenStore].
func (p *PersistentRefreshTokenStore) FindByHash(ctx context.Context, hash string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if result := db.GetDB(ctx).Where("hash = ?", hash).First(&refreshToken); result.Error != nil {
		return nil, result.Error
	}
	return &refreshToken, nil
}

// Revoke implements [RefreshTokenStore].
func (p *PersistentRefreshTokenStore) Revoke(ctx context.Context, id uuid.UUID) error {
	result := db.GetDB(ctx).
		Model(&models.RefreshToken{}).
		Where("id = ?", id.String()).
		Update("revoked", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PersistentRefreshTokenStore) RevokeAll(ctx context.Context, userID uuid.UUID, clientID string) error {
	result := db.GetDB(ctx).
		Model(&models.RefreshToken{}).
		Where("identity_id = ? AND client_id = ?", userID.String(), clientID).
		Update("revoked", true)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// Save implements [RefreshTokenStore].
func (p *PersistentRefreshTokenStore) Save(ctx context.Context, token *models.RefreshToken) error {
	if result := db.GetDB(ctx).Create(&token); result.Error != nil {
		return result.Error
	}

	return nil
}
