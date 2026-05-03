package email

import (
	"context"
	"encoding/json"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
)

type EmailPayload struct {
	To      string
	From    string
	Subject string
	Body    string
	// TODO:
	// - email type - to get proper template
	// - then instead of body - use data
	// - sender will use type and data to get email template from registry, build it with data and send
	// - use templ for templates
}

func ProvideSendEmailJobHandler(config *config.Config) *SendEmailJob {
	return &SendEmailJob{
		emailService: &EmailService{config: config},
	}
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
