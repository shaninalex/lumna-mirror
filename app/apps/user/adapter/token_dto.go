package adapter

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/internal/db"
)

// UserTokenDto is a Data Transfer Object (DTO) for user tokens.
// It is used to expose token information to the client (e.g., via an API)
// while hiding sensitive fields like the actual refresh token string.
type UserTokenDto struct {
	Id               int64      `json:"id"`                 // Unique ID of the token record
	Device           string     `json:"device"`             // Device or client info associated with the token
	RefreshExpiresAt time.Time  `json:"refresh_expires_at"` // Expiration time of the refresh token
	Revoked          bool       `json:"revoked"`            // Flag indicating if the token was revoked
	RevokedAt        *time.Time `json:"revoked_at"`         // Timestamp when the token was revoked (nil if not revoked)
	CreatedAt        time.Time  `json:"created_at"`         // Timestamp when the token was created
}

// ToUserTokenDto converts a models.UserToken entity from the database
// into a UserTokenDto for safe exposure to clients.
func ToUserTokenDto(token *db.UserToken) *UserTokenDto {
	return &UserTokenDto{
		Id:               token.ID,
		Device:           token.Device,
		RefreshExpiresAt: token.RefreshExpiresAt,
		Revoked:          token.Revoked,
		RevokedAt:        token.RevokedAt,
		CreatedAt:        token.CreatedAt,
	}
}

// ToUserTokenDtoList converts a slice of UserToken entities into a slice of UserTokenDto.
// Useful for returning a list of user tokens in an API response.
func ToUserTokenDtoList(tokens []*db.UserToken) []*UserTokenDto {
	output := make([]*UserTokenDto, len(tokens)) // pre-allocate slice for performance
	for i, token := range tokens {
		output[i] = ToUserTokenDto(token) // convert each token individually
	}
	return output
}
