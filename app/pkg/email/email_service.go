package email

import (
	"context"
)

type EmailSender interface {
	Send(ctx context.Context, to, from, subject, html string) error
}

type EmailService struct {
}

func (s *EmailService) Send(ctx context.Context, to, from, subject, html string) error {
	panic("implement me")
}
