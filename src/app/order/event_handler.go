package order

import (
	"context"
	"fmt"
	"github.com/mjedari/go-cqrs-template/app/coin"
	coinDomain "github.com/mjedari/go-cqrs-template/domain/coin"
	"github.com/mjedari/go-cqrs-template/domain/order"
	"github.com/mjedari/go-cqrs-template/infra/providers/messaging"
	"github.com/sirupsen/logrus"
)

type OrderEventHandler struct {
	orderRepository *OrderRepository
	coinRepository  *coin.CoinRepository
	bus             *messaging.EventBus
}

func NewOrderEventHandler(orderRepository *OrderRepository, coinRepository *coin.CoinRepository, bus *messaging.EventBus) *OrderEventHandler {
	return &OrderEventHandler{orderRepository: orderRepository, coinRepository: coinRepository, bus: bus}
}

func (h OrderEventHandler) HandleEvents(ctx context.Context, events []interface{}) error {

	for _, event := range events {
		switch event.(type) {
		case order.InstantOrderEvent:
			instantEvent := event.(order.InstantOrderEvent)
			h.handleInstantBuyEvent(ctx, instantEvent)
		case order.OrderEvent:
			orderEvent := event.(order.OrderEvent)
			h.handleBuyEvent(ctx, orderEvent)
		case order.FailTransactionEvent:
			failEvent := event.(order.FailTransactionEvent)
			h.handleFailTransactionEvent(ctx, failEvent)
		case order.TestEvent:
			testEvent := event.(order.TestEvent)
			h.handleTestEvent(ctx, testEvent)
		}
	}
	return nil
}

func (h OrderEventHandler) handleBuyEvent(ctx context.Context, orderEvent order.OrderEvent) {
	fmt.Println("handle order event", orderEvent)

}

func (h OrderEventHandler) handleInstantBuyEvent(ctx context.Context, orderEvent order.InstantOrderEvent) {
	fmt.Println("handle order event", orderEvent)
	coin := coinDomain.Coin{Id: orderEvent.CoinId}
	if err := h.coinRepository.GetCoin(ctx, &coin); err != nil {
		//todo: handle event error
		logrus.Error("Error")
	}

	initOrders, err := h.orderRepository.GetInitializedOrders(ctx)
	if err != nil {
		// todo: handle err
		logrus.Error(err)
	}

	transaction := ExchangeTransaction{
		Coin:     coin,
		Quantity: initOrders.GetTotalQuantity(),
	}

	if exErr := h.buyFromExchange(ctx, transaction); exErr != nil {
		logrus.Error("transaction failed")
		event := order.FailTransactionEvent{
			Orders: initOrders.List,
		}
		if transactionFailedErr := h.bus.Publish(ctx, []interface{}{event}); transactionFailedErr != nil {

		}
	}

	if err := h.orderRepository.SettleOrders(ctx, initOrders.List); err != nil {

	}
}

func (h OrderEventHandler) handleFailTransactionEvent(ctx context.Context, event order.FailTransactionEvent) {
	for _, failedOrder := range event.Orders {
		failedOrder.ChangeStatus(order.FAILED)
		if err := h.orderRepository.UpdateOrder(ctx, &failedOrder); err != nil {
			// todo: handle
		}
	}

}

func (h OrderEventHandler) handleTestEvent(ctx context.Context, testEvent order.TestEvent) {
	fmt.Println("handle order event", testEvent)

}

func (h OrderEventHandler) buyFromExchange(ctx context.Context, transaction ExchangeTransaction) error {
	// call exchange api to buy
	fmt.Println("called external request")
	return nil
}
