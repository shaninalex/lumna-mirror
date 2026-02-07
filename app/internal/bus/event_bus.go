package bus

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal"
)

type EventHandlerFunc func(ctx context.Context, data any)

type EventBus interface {
	Subscribe(event internal.Event, callback EventHandlerFunc)
	Publish(ctx context.Context, event internal.Event, data any)
}

var _ EventBus = (*bus)(nil)

type bus struct {
	events map[internal.Event][]EventHandlerFunc
}

func ProvideEventBus() EventBus {
	s := &bus{
		events: make(map[internal.Event][]EventHandlerFunc),
	}
	return s
}

func (s *bus) Subscribe(event internal.Event, callback EventHandlerFunc) {
	s.events[event] = append(s.events[event], callback)
}

func (s *bus) Publish(ctx context.Context, event internal.Event, data any) {
	subscribers, ok := s.events[event]
	if !ok {
		return
	}

	for _, subscriber := range subscribers {
		subscriber(ctx, data)
	}
}
