package observer

import (
	"context"
	"fmt"
)

type Event string

type EventHandlerFunc func(ctx context.Context, data any)

type Observer interface {
	Subscribe(event Event, callback EventHandlerFunc)
	Publish(ctx context.Context, event Event, data any)
}

var _ Observer = (*observer)(nil)

type observer struct {
	events map[Event][]EventHandlerFunc
}

func ProvideEventBus() Observer {
	s := &observer{
		events: make(map[Event][]EventHandlerFunc),
	}
	return s
}

func (s *observer) Subscribe(event Event, callback EventHandlerFunc) {
	s.events[event] = append(s.events[event], callback)
}

func (s *observer) Publish(ctx context.Context, event Event, data any) {
	subscribers, ok := s.events[event]
	if !ok {
		return
	}

	for _, subscriber := range subscribers {
		fmt.Println("[BUS] publish ", event)
		subscriber(ctx, data)
	}
}
