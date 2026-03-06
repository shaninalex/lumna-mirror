package email

import (
	"context"
)

type EmailService struct {
}

func (s *EmailService) Send(ctx context.Context, to, from, subject, html string) error {
	//TODO implement me
	panic("implement me")
}
