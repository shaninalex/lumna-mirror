package email

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/repositories"
)

const (
	workersNumber = 2
)

type EmailQueue struct {
	workers    uint8
	repository repositories.EmailRepository
}

func NewEmailQueue(repository repositories.EmailRepository) *EmailQueue {
	return &EmailQueue{
		workers:    workersNumber,
		repository: repository,
	}
}

// Process run the queue, ask db each tick and process pending emails
func (s *EmailQueue) Process(ctx context.Context) {

}
