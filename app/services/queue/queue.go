package queue

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/pkg/email"
	"gitlab.com/shaninalex/lumna/app/pkg/logger"
	"gorm.io/gorm"
)

func ProvideJobQueueService(
	ctx context.Context,
	db *gorm.DB,
	logger logger.Logger,
	config *config.Config,
) JobQueueService {
	s := JobQueueService{
		ctx:      ctx,
		db:       db,
		logger:   logger,
		handlers: make(map[string]JobHandler),
	}

	s.init(config)
	s.run()
	return s
}

type JobQueueService struct {
	db       *gorm.DB
	ctx      context.Context
	logger   logger.Logger
	handlers map[string]JobHandler
}

func (s *JobQueueService) init(config *config.Config) {
	// TODO: how to provide dependencies for handlers? Via init args - that sucks... Need to build dig.Container properly
	s.Register("send_email", email.ProvideSendEmailJobHandler(config))
}

func (s *JobQueueService) run() {
	jobChan := make(chan models.Job, 10)

	for i := 0; i < 5; i++ {
		go func() {
			for {
				select {
				case <-s.ctx.Done():
					fmt.Println("[JobQueueService] shutdown")
					return
				case job := <-jobChan:
					s.process(job)
				}
			}
		}()
	}

	go func() {
		for {
			jobs, err := s.getJobs()
			if err != nil {
				s.logger.Log(fmt.Sprintf("[JobQueueService] get jobs error: %v", err))
				break
			}
			s.logger.Log("[JobQueueService] tick...")

			for _, job := range jobs {
				jobChan <- job
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

func (s *JobQueueService) getJobs() ([]models.Job, error) {
	jobs := []models.Job{}
	if result := s.db.WithContext(s.ctx).Where("status = ?", models.JobStatusPending).Limit(5).Find(&jobs); result.Error != nil {
		return nil, result.Error
	}
	return jobs, nil
}

func (s *JobQueueService) process(job models.Job) {
	ctx, cancel := context.WithTimeout(s.ctx, 15*time.Second)
	defer cancel()

	attempts := job.Attempts
	attempts += 1
	handler, ok := s.handlers[job.Type]
	if !ok {
		s.logger.Log(fmt.Sprintf("unknown job type: %s", job.Type))
		return
	}

	err := handler.Handle(ctx, &job)
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
