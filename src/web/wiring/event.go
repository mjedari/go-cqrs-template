package wiring

import (
	"github.com/mjedari/go-cqrs-template/infra/providers/messaging"
)

func (w *Wire) GetEventBus() *messaging.EventBus {
	if w.eventBus == nil {
		w.eventBus = messaging.NewEventBus("my-event-bus", w.GetRedis())
	}

	return w.eventBus
}
