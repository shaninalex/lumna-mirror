package email

import (
	"context"
	"encoding/json"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/bus"
	"gitlab.com/shaninalex/lumna/app/internal/logger"
)

type emailQueue struct {
	bus    bus.EventBus
	ch     chan any // TODO: make email queue type
	ctx    context.Context
	logger logger.Logger
}

func InvokeEmailQueue(
	bus bus.EventBus,
	logger logger.Logger,
	ctx context.Context,
) error {
	s := &emailQueue{
		bus:    bus,
		ch:     make(chan any, 100),
		ctx:    ctx,
		logger: logger,
	}
	s.init()
	return nil
}

func (s *emailQueue) addToEmailQueue(ctx context.Context, data any) {
	s.ch <- data
}

func (s *emailQueue) init() {
	s.bus.Subscribe(internal.EmailSendEvent, s.addToEmailQueue)
	go func() {
		fmt.Println("[EmailQueue] Start")
		for {
			select {
			case <-s.ctx.Done():
				// Process current and stop
				fmt.Println("[EmailQueue] Stopped")
				return

			case emailInQueue := <-s.ch:
				b, _ := json.Marshal(emailInQueue)
				fmt.Println("Email queue processing the data", string(b))
			}
		}
	}()
}
