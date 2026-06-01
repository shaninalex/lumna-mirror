package invitation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/observer"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_InvitationService_Create(t *testing.T) {
	db := testutils.ProvideTestDB()
	_ = testutils.Migrate(db)
	s := ProvideInvitationService(
		repositories.NewGormInvitationRepository(db),
		repositories.NewGormEmailRepository(db),
		observer.ProvideEventBus(),
	)
	ctx := context.Background()

	defer testutils.ClearDB(db)

	invitation, emailLink, err := s.Create(ctx, "test@test.com", "user", nil)
	assert.NoError(t, err)
	assert.NotNil(t, invitation)
	assert.NotEmpty(t, emailLink)
	assert.Equal(t, utils.HashToken(emailLink), invitation.TokenHash)
	assert.Equal(t, invitation.State, models.InvitationStatePending)
	assert.False(t, invitation.IsValid())
	assert.False(t, invitation.IsExpired())

	invitation.State = models.InvitationStateSent
	assert.True(t, invitation.IsValid())

	invitations, err := s.List(ctx)
	assert.NoError(t, err)
	assert.Len(t, invitations, 1)
}
