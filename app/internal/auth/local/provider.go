package local

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"gitlab.com/shaninalex/lumna/app/internal/auth"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

// var _ auth.AuthProvider = (*LocalAuthProvider)(nil)

type PasswordCredentials struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (s *PasswordCredentials) Validate() error {
	if _, err := mail.ParseAddress(s.Email); err != nil {
		return err
	}

	if len(s.Password) == 0 {
		return fmt.Errorf("password is empty")
	}

	return nil
}

type LocalAuthProvider struct {
}

func NewLocalAuthProvider() *LocalAuthProvider {
	s := &LocalAuthProvider{}
	return s
}

func (s *LocalAuthProvider) Name() string {
	return "local"
}

var UserNotFoundError = errors.New("user not found")

func (s *LocalAuthProvider) Authenticate(ctx context.Context, payload *PasswordCredentials) (*auth.AuthResult, error) {
	database := db.GetDB(ctx)
	credentials := models.Credential{}
	if result := database.Preload("Identity").First(&credentials, "email = ?", payload.Email); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, UserNotFoundError
		}
		return nil, result.Error
	}

	if credentials.PasswordHash == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := ValidatePassword(*credentials.PasswordHash, payload.Password); err != nil {
		return nil, err
	}

	return &auth.AuthResult{
		Identity:   &credentials.Identity,
		Provider:   "local",
		Credential: &credentials,
	}, nil
}
