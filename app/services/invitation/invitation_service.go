package invitation

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/email"
	"gitlab.com/shaninalex/lumna/app/services/email/templates"
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

type InvitationEmailMeta struct {
	InvitationID uint `json:"invitation_id"`
}

var (
	defaultValidUntil = 1 * time.Hour
)

type Service interface {
	Create(ctx context.Context, email, role string, meta map[string]any) (*models.Invitation, string, error) // TODO: no need to return token string
	Get(ctx context.Context, hash string) (*models.Invitation, error)                                        // TODO: rename into GetByHash
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
) Service {
	s := &service{
		repository:      repository,
		emailRepository: emailRepository,
		bus:             bus,
	}
	s.init()
	return s
}

type service struct {
	repository      repositories.InvitationRepository
	emailRepository repositories.EmailRepository
	bus             observer.Observer
}

var _ Service = (*service)(nil)

func (s *service) init() {
	s.bus.Subscribe(email.EventEmailQueueSent, s.handleEventInvitationEmailSent)
}

func (s *service) List(ctx context.Context) ([]models.Invitation, error) {
	return s.repository.List(ctx)
}

func (s *service) Create(ctx context.Context, email, role string, meta map[string]any) (*models.Invitation, string, error) {
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

func (s *service) Accept(ctx context.Context, token string) error {
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

func (s *service) Validate(ctx context.Context, token string) error {
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

func (s *service) Get(ctx context.Context, token string) (*models.Invitation, error) {
	invitation, err := s.repository.GetByHash(ctx, utils.HashToken(token))
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func (s *service) Delete(ctx context.Context, invitationId uint) error {
	return s.repository.Delete(ctx, invitationId)
}

func (s *service) Reset(ctx context.Context, invitationId uint) (string, error) {
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

func (s *service) handleEventInvitationEmailSent(ctx context.Context, data any) {
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

type OnboardingCreate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
