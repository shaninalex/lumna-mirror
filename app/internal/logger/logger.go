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

func NewSimpleLogger(ctx context.Context) *SimpleLogger {
	l := SimpleLogger{
		queue: make(chan string, 100),
		ctx:   ctx,
	}
	l.init()
	return &l
}

func (s *SimpleLogger) Log(msg string) {
	select {
	case s.queue <- msg:
	// case <-s.ctx.Done():
	// fmt.Println("[Logger] log queue closed ")
	default:
		fmt.Println("logger queue is full...")
	}
}

func (s *SimpleLogger) init() {
	go func() {
		for {
			select {
			case msg := <-s.queue:
				s.process(msg)
			case <-s.ctx.Done():
				s.drain()
				return
			}
		}
	}()
}

func (s *SimpleLogger) drain() {
	fmt.Println("[Logger] remaining logs...")
	for {
		select {
		case msg := <-s.queue:
			s.process(msg)
		default:
			return
		}
	}
}

func (s *SimpleLogger) process(msg string) {
	log.Println(msg)
}
