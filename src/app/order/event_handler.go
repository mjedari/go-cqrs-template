package order

import (
	"context"
	"fmt"
	"github.com/mjedari/go-cqrs-template/domain/order"
	"github.com/mjedari/go-cqrs-template/infra/providers/messaging"
)

type OrderEventHandler struct {
	bus *messaging.EventBus
}

func NewOrderEventHandler(bus *messaging.EventBus) *OrderEventHandler {
	return &OrderEventHandler{bus: bus}
}

func (h OrderEventHandler) HandleEvents(ctx context.Context, events []interface{}) error {
	for _, event := range events {
		switch event.(type) {
		case order.OrderEvent:
			e := event.(order.OrderEvent)
			h.handleBuyEvent(ctx, e)
		case order.TestEvent:
			e := event.(order.TestEvent)
			h.HandleTestEvent(ctx, e)
		}
	}

	return nil
}

func (h OrderEventHandler) handleBuyEvent(ctx context.Context, orderEvent order.OrderEvent) {
	fmt.Println("handle order event", orderEvent)
}

func (h OrderEventHandler) HandleTestEvent(ctx context.Context, testEvent order.TestEvent) {
	fmt.Println("handle order event", testEvent)

}
