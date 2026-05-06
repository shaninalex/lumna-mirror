package auth

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type PasswordCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (s *PasswordCredentials) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return nil
}

type EmailAuthProvider struct {
	credentialRepository repositories.CredentialRepository
	identityRepository   repositories.IdentityRepository
}

func NewEmailAuthProvider(
	credentialRepository repositories.CredentialRepository,
	identityRepository repositories.IdentityRepository,
) *EmailAuthProvider {
	s := &EmailAuthProvider{
		credentialRepository: credentialRepository,
		identityRepository:   identityRepository,
	}
	return s
}

func (s *EmailAuthProvider) Name() string {
	return "local"
}

func (s *EmailAuthProvider) Authenticate(ctx context.Context, payload *PasswordCredentials) (*models.Identity, error) {
	credential, err := s.credentialRepository.GetByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	if credential.PasswordHash == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := ValidatePassword(*credential.PasswordHash, payload.Password); err != nil {
		return nil, err
	}

	identity, err := s.identityRepository.GetIdentityByID(ctx, credential.IdentityID)
	if err != nil {
		return nil, err
	}

	return identity, nil
}
