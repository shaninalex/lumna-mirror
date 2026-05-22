package services

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/observer"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/email"
	"gitlab.com/shaninalex/lumna/app/services/email/templates"
)

type InvitationEmailMeta struct {
	InvitationID uint `json:"invitation_id"`
}

var (
	defaultValidUntil = 1 * time.Hour
)

type InvitationManager interface {
	Create(ctx context.Context, email, role string, meta map[string]any) (*models.Invitation, string, error)
	Get(ctx context.Context, hash string) (*models.Invitation, error) // TODO: rename into GetByHash
	Accept(ctx context.Context, token string) error
	Validate(ctx context.Context, token string) error
	Delete(ctx context.Context, invitationId uint) error
	Reset(ctx context.Context, invitationId uint) (string, error)
	List(ctx context.Context) ([]models.Invitation, error)
}

func ProvideInvitationService(
	repository repositories.InvitationRepository,
	emailRepository repositories.EmailRepository,
	bus observer.Observer,
) InvitationManager {
	s := &InvitationService{
		repository:      repository,
		emailRepository: emailRepository,
		bus:             bus,
	}
	s.init()
	return s
}

type InvitationService struct {
	repository      repositories.InvitationRepository
	emailRepository repositories.EmailRepository
	bus             observer.Observer
}

func (s *InvitationService) init() {
	s.bus.Subscribe(email.EventEmailQueueSent, s.handleEventInvitationEmailSent)
}

func (s *InvitationService) List(ctx context.Context) ([]models.Invitation, error) {
	return s.repository.List(ctx)
}

func (s *InvitationService) Create(ctx context.Context, email, role string, meta map[string]any) (*models.Invitation, string, error) {
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
	if meta != nil {
		invitation.Meta = meta
	}
	if err = s.repository.Create(ctx, invitation); err != nil {
		return nil, "", err
	}

	t := templates.NewEmailInvitationEmailTemplate(
		invitation.Email,
		fmt.Sprintf("http://localhost:4200/auth/accept-invite/%s", token), // Build server url
	)

	metaMap, err := models.MetaToMap(InvitationEmailMeta{InvitationID: invitation.ID})
	if err != nil {
		return nil, "", err
	}

	eml := &models.Email{
		ToEmail:   invitation.Email,
		FromEmail: "your@server.host", // from settings somewhere ?
		Body:      t.HTML(),
		Subject:   "You have been invited",
		Meta:      metaMap,
	}

	if err = s.emailRepository.Create(ctx, eml); err != nil {
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

func (s *InvitationService) Validate(ctx context.Context, token string) error {
	tokenHash := utils.HashToken(token)
	invitation, err := s.repository.GetByHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	if !invitation.IsValid() {
		return errors.New("invitation invalid or expired")
	}

	invitation.State = models.InvitationStateValidated
	return s.repository.Update(ctx, invitation)
}

func (s *InvitationService) Get(ctx context.Context, token string) (*models.Invitation, error) {
	invitation, err := s.repository.GetByHash(ctx, utils.HashToken(token))
	if err != nil {
		return nil, err
	}
	return invitation, nil
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

func (s *InvitationService) handleEventInvitationEmailSent(ctx context.Context, data any) {
	m, ok := data.(models.MetaContaing)
	if !ok {
		return
	}

	var meta InvitationEmailMeta
	if err := models.MapToMeta(m.GetMeta(), &meta); err != nil {
		return
	}
	if meta.InvitationID == 0 {
		return
	}

	inv, err := s.repository.GetById(ctx, meta.InvitationID)
	if err != nil {
		return
	}

	inv.State = models.InvitationStateSent
	if err = s.repository.Update(ctx, inv); err != nil {
		return
	}
}
