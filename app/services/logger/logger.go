package logger

import (
	"context"
	"fmt"
	"log"
)

type Logger interface {
	Log(msg string)
}

type simpleLogger struct {
	ctx   context.Context
	queue chan string
}

func ProvideLogger(ctx context.Context) Logger {
	l := simpleLogger{
		queue: make(chan string, 100),
		ctx:   ctx,
	}
	l.init()
	return &l
}

func (s *simpleLogger) Log(msg string) {
	s.queue <- msg
}

func (s *simpleLogger) init() {
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

func (s *simpleLogger) process(msg string) {
	log.Println(msg)
}
