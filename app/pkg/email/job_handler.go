package email

import (
	"context"
	"encoding/json"

	"gitlab.com/shaninalex/lumna/app/models"
)

type EmailSender interface {
	Send(ctx context.Context, to, from, subject, html string) error
}

type EmailPayload struct {
	To      string
	From    string
	Subject string
	Body    string
}

func ProvideSendEmailJob() *SendEmailJob {
	return &SendEmailJob{}
}

type SendEmailJob struct {
	emailService EmailSender
}

func (j *SendEmailJob) Handle(ctx context.Context, job *models.Job) error {
	payload := EmailPayload{}

	err := json.Unmarshal([]byte(job.Payload), &payload)
	if err != nil {
		return err
	}

	return j.emailService.Send(ctx, payload.To, payload.From, payload.Subject, payload.Body)
}
