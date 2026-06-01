package logger

import (
	"context"
	"fmt"
	"log"
)

type Logger interface {
	Log(msg string)
}

type SimpleLogger struct {
	ctx   context.Context
	queue chan string
}

func ProvideLogger(ctx context.Context) Logger {
	l := SimpleLogger{
		queue: make(chan string, 100),
		ctx:   ctx,
	}
	l.init()
	return &l
}

func (s *SimpleLogger) Log(msg string) {
	s.queue <- msg
}

func (s *SimpleLogger) init() {
	fmt.Println("[Logger] Start")
	go func() {
		for {
			select {
			case msg := <-s.queue:
				s.process(msg)
			case <-s.ctx.Done():
				fmt.Println("[Logger] Stopped")
				return
			}
		}
	}()
}

func (s *SimpleLogger) process(msg string) {
	log.Println(msg)
}
