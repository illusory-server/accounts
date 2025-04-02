package broker

import (
	"context"
	"time"

	"github.com/illusory-server/accounts/pkg/logger"
)

type Event struct {
	ID        string
	AccountID string
	Type      string
	Data      any
	Timestamp time.Time
}

type Eventer interface {
	GetAll(ctx context.Context) ([]Event, error)
	RemoveById(ctx context.Context, id string) error
}

type Sender interface {
	Send(ctx context.Context, e Event) error
}

type Handler struct {
	event Eventer
	send  Sender
	log   logger.Logger
}

func (h *Handler) Handle(ctx context.Context) error {
	events, err := h.event.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, event := range events {
		err := h.send.Send(ctx, event)
		if err != nil {
			h.log.Warn(ctx, "send event failed",
				logger.String("event_id", event.ID),
				logger.Any("event", event))
			continue
		}
		err = h.event.RemoveById(ctx, event.ID)
		if err != nil {
			h.log.Error(ctx, "remove event by id failed",
				logger.String("event_id", event.ID))
		}
	}

	return nil
}
