package email

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/pkg/config"
	gomail "gopkg.in/gomail.v2"
)

type EmailSender interface {
	Send(ctx context.Context, to, from, subject, html string) error
}

func ProvideEmailService(config *config.Config) EmailSender {
	return &emailService{
		config: config,
	}
}

type emailService struct {
	config *config.Config
}

func (s *emailService) Send(ctx context.Context, to, from, subject, html string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", html)

	dialer := gomail.NewDialer(
		s.config.String("email.host"),
		s.config.Int("email.port"),
		s.config.String("email.username"),
		s.config.String("email.password"),
	)

	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}

	return nil
}
