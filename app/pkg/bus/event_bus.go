package bus

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/pkg"
)

type EventHandlerFunc func(ctx context.Context, data any)

type EventBus interface {
	Subscribe(event pkg.Event, callback EventHandlerFunc)
	Publish(ctx context.Context, event pkg.Event, data any)
}

var _ EventBus = (*bus)(nil)

type bus struct {
	events map[pkg.Event][]EventHandlerFunc
}

func ProvideEventBus() EventBus {
	s := &bus{
		events: make(map[pkg.Event][]EventHandlerFunc),
	}
	return s
}

func (s *bus) Subscribe(event pkg.Event, callback EventHandlerFunc) {
	s.events[event] = append(s.events[event], callback)
}

func (s *bus) Publish(ctx context.Context, event pkg.Event, data any) {
	subscribers, ok := s.events[event]
	if !ok {
		return
	}

	for _, subscriber := range subscribers {
		subscriber(ctx, data)
	}
}
