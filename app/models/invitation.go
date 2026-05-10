package models

import (
	"time"

	"gorm.io/gorm"
)

type InvitationState string

var (
	// InvitationStatePending just created
	InvitationStatePending = InvitationState("pending")

	// InvitationStateSent when invitation is being sent via email ( or other transport )
	InvitationStateSent = InvitationState("sent")

	// InvitationStateAccepted when user confirms invitation successfully
	InvitationStateAccepted = InvitationState("accepted")

	// InvitationStateRevoked manually revoked and invitation receiver unable to use that inv. link anymore
	InvitationStateRevoked = InvitationState("revoked")
)

type Invitation struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Email       string          `gorm:"unique" json:"email"`
	TokenHash   string          `gorm:"uniqueIndex" json:"-"`
	State       InvitationState `json:"state"`
	Role        string          `json:"role"`
	WorkspaceID uint            `gorm:"not null;index" json:"workspace_id"`
	CreatedAt   time.Time       `gorm:"not null" json:"created_at"`
	ValidUntil  time.Time       `gorm:"not null" json:"valid_until"`
}

// BeforeCreate set's created at time
func (s *Invitation) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *Invitation) IsExpired() bool {
	return time.Now().After(s.ValidUntil)
}

// IsValid - check is it expired or not sent.
// Not sent mean's that in any other states it's not valid invitation
// pending - it's not sent to receiver and not valid to use yet
// accepted - you can't reuse accepted invitation
// revoked - it's manually revoked by admin and can't be used
func (s *Invitation) IsValid() bool {
	return s.State == InvitationStateSent && !s.IsExpired()
}

// Reset - resets Invitation, set new ValidUntil and make it Used=false
func (s *Invitation) Reset(newReset time.Time, newTokenHash string) {
	s.ValidUntil = newReset
	s.State = InvitationStatePending
	s.TokenHash = newTokenHash
}
