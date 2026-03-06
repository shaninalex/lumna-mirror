package queue

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/email"
	"gitlab.com/shaninalex/lumna/app/pkg/logger"
	"gorm.io/gorm"
)

func ProvideJobQueueService(
	ctx context.Context,
	db *gorm.DB,
	logger logger.Logger,
) JobQueueService {
	s := JobQueueService{
		ctx:      ctx,
		db:       db,
		logger:   logger,
		handlers: make(map[string]JobHandler),
	}
	s.init()
	s.run()
	return s
}

type JobQueueService struct {
	db       *gorm.DB
	ctx      context.Context
	logger   logger.Logger
	handlers map[string]JobHandler
}

func (s *JobQueueService) init() {

	// handlers require dependencies!
	s.Register("send_email", email.ProvideSendEmailJob())
}

func (s *JobQueueService) run() {
	go func() {
		for {
			jobs := []models.Job{}
			result := s.db.WithContext(s.ctx).Where("status = ?", models.JobStatusPending).Find(&jobs)
			if result.Error != nil {
				// restart?
				s.logger.Log(fmt.Sprintf("[JobQueueService] get jobs error: %v", result.Error))
				break
			}

			for _, job := range jobs {
				attempts := job.Attempts
				attempts += 1
				handler, ok := s.handlers[job.Type]
				if !ok {
					s.logger.Log(fmt.Sprintf("unknown job type: %s", job.Type))
					continue
				}

				err := handler.Handle(s.ctx, &job)
				if err != nil {
					if attempts == 3 {
						job.Status = models.JobStatusError
					} else {
						job.Status = models.JobStatusRepeat
					}
					s.logger.Log(fmt.Sprintf("[JobQueueService] handle job error: %v", err))
				} else {
					job.Status = models.JobStatusSuccess
				}
				s.db.Save(&job)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
