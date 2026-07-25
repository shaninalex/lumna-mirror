package observer

import (
	"context"
	"fmt"
	"sync"
)

type Event string

type EventHandlerFunc func(ctx context.Context, data any)

type Observer interface {
	Subscribe(event Event, callback EventHandlerFunc)
	Publish(ctx context.Context, event Event, data any)
}

var _ Observer = (*observer)(nil)

type observer struct {
	mu     sync.RWMutex
	events map[Event][]EventHandlerFunc
}

func ProvideEventBus() Observer {
	s := &observer{
		events: make(map[Event][]EventHandlerFunc),
	}
	return s
}

func (s *observer) Subscribe(event Event, callback EventHandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event] = append(s.events[event], callback)
}

func (s *observer) Publish(ctx context.Context, event Event, data any) {
	s.mu.RLock()
	subscribers := s.events[event]
	s.mu.RUnlock()
	if len(subscribers) == 0 {
		return
	}

	fmt.Println("[BUS] publish ", event)
	// Subscribers run asynchronously, so they must not be tied to the
	// caller's request-scoped context: once the request completes its ctx is
	// canceled, which would abort in-flight subscriber work (e.g. DB queries
	// failing with "context canceled"). WithoutCancel keeps the context values
	// while detaching cancellation/deadline propagation.
	subCtx := context.WithoutCancel(ctx)
	for _, subscriber := range subscribers {
		go subscriber(subCtx, data)
	}
}
