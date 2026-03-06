package queue

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
)

type JobHandler interface {
	Handle(ctx context.Context, job *models.Job) error
}

func (s *JobQueueService) Register(jobType string, handler JobHandler) {
	s.handlers[jobType] = handler
}

// NOTE: will be great if some service is unavailable - deregister handlers
