package email

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

const (
	workersNumber                      = 2
	queueBuffer                        = 100
	tickInterval                       = 1 * time.Second
	EventEmailQueueSent observer.Event = "EMAIL_QUEUE_SENT"
)

type EmailQueue struct {
	ctx        context.Context
	workers    int
	repository repositories.EmailRepository
	sender     EmailSender
	jobs       chan models.Email
	bus        observer.Observer
}

func NewEmailQueue(
	ctx context.Context,
	repository repositories.EmailRepository,
	sender EmailSender,
	bus observer.Observer,
) *EmailQueue {
	s := &EmailQueue{
		ctx:        ctx,
		workers:    workersNumber,
		repository: repository,
		sender:     sender,
		jobs:       make(chan models.Email, queueBuffer),
		bus:        bus,
	}

	s.process()
	return s
}

// Process run the queue, ask db each tick and process pending emails
func (s *EmailQueue) process() {
	fmt.Println("[EmailQueue] Start")
	for i := 0; i < s.workers; i++ {
		go s.worker(i)
	}

	go s.processQueue()
}

func (s *EmailQueue) worker(id int) {
	for {
		select {
		case <-s.ctx.Done():
			fmt.Printf("[EmailQueue] worker %d stopped\n", id)
			return
		case email := <-s.jobs:
			s.handle(email)
		}
	}
}

func (s *EmailQueue) handle(email models.Email) {
	err := s.sender.Send(s.ctx, email.ToEmail, email.FromEmail, email.Subject, email.Body)
	now := time.Now()
	if err != nil {
		fmt.Printf("[EmailQueue] send error id=%d: %v\n", email.ID, err)
		email.Status = models.EmailStatusError
	} else {
		email.Status = models.EmailStatusSuccess
		email.SentAt = &now
	}

	if err := s.repository.Update(s.ctx, &email); err != nil {
		fmt.Printf("[EmailQueue] update error id=%d: %v\n", email.ID, err)
	}

	s.bus.Publish(s.ctx, EventEmailQueueSent, email)
}

func (s *EmailQueue) processQueue() {
	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			fmt.Println("[EmailQueue] dispatcher stopped")
			return
		case <-ticker.C:
			s.dispatch()
		}
	}
}

func (s *EmailQueue) dispatch() {
	pending := s.repository.ListPending(s.ctx)
	for _, email := range pending {
		email.Status = models.EmailStatusRunning
		if err := s.repository.Update(s.ctx, &email); err != nil {
			fmt.Printf("[EmailQueue] mark running error id=%d: %v\n", email.ID, err)
			continue
		}

		select {
		case <-s.ctx.Done():
			return
		case s.jobs <- email:
		}
	}
}
