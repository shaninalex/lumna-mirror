package services

import (
	"context"
	"errors"
	"net/mail"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

var (
	defaultValidUntil = 1 * time.Hour
)

type InvitationManager interface {
	Create(ctx context.Context, email, role string) (*models.Invitation, string, error)
	Accept(ctx context.Context, token string) error
	Delete(ctx context.Context, invitationId uint) error
	Reset(ctx context.Context, invitationId uint) (string, error)
	List(ctx context.Context) ([]models.Invitation, error)
}

func ProvideInvitationService(repository repositories.InvitationRepository) InvitationManager {
	return &InvitationService{repository: repository}
}

type InvitationService struct {
	repository repositories.InvitationRepository
}

func (s *InvitationService) List(ctx context.Context) ([]models.Invitation, error) {
	return s.repository.List(ctx)
}

func (s *InvitationService) Create(ctx context.Context, email, role string) (*models.Invitation, string, error) {
	if email == "" || role == "" {
		return nil, "", errors.New("invalid data")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, "", err
	}

	// NOTE:
	// token -> send in email,
	// tokenHash -> save in db
	token, tokenHash, err := utils.GenerateToken()
	if err != nil {
		return nil, "", err
	}
	invitation := &models.Invitation{
		Email:      email,
		Role:       role,
		State:      models.InvitationStatePending,
		TokenHash:  tokenHash,
		ValidUntil: time.Now().Add(defaultValidUntil),
	}
	if err = s.repository.Create(ctx, invitation); err != nil {
		return nil, "", err
	}
	return invitation, token, nil
}

func (s *InvitationService) Accept(ctx context.Context, token string) error {
	tokenHash := utils.HashToken(token)

	invitation, err := s.repository.GetByHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	if !invitation.IsValid() {
		return errors.New("invitation invalid or expired")
	}

	invitation.State = models.InvitationStateAccepted
	return s.repository.Update(ctx, invitation)
}

func (s *InvitationService) Delete(ctx context.Context, invitationId uint) error {
	return s.repository.Delete(ctx, invitationId)
}

func (s *InvitationService) Reset(ctx context.Context, invitationId uint) (string, error) {
	invitation, err := s.repository.GetById(ctx, invitationId)
	if err != nil {
		return "", err
	}

	token, tokenHash, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}

	invitation.Reset(time.Now().Add(defaultValidUntil), tokenHash)
	if err := s.repository.Update(ctx, invitation); err != nil {
		return "", err
	}

	return token, nil
}
