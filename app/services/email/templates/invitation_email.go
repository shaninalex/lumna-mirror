package templates

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/utils"
)

type EmailInvitation struct {
	email string

	// This is full confirmation link. EmailInvitation should not know about
	// token, domains, configurations etc. It's just email template. It should
	// cary about template, styles and format of given data.
	link string
}

func NewEmailInvitationEmailTemplate(email, link string) *EmailInvitation {
	return &EmailInvitation{
		email: email,
		link:  link,
	}
}

func (s *EmailInvitation) receiverEmail() string {
	return s.email
}

func (s *EmailInvitation) receiverConfirmationLink() string {
	return s.link
}

func (s *EmailInvitation) HTML() string {
	return utils.TemplToString(context.Background(), s.template())
}
