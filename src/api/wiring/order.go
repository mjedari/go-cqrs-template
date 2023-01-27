package wiring

import (
	order2 "github.com/mjedari/go-cqrs-template/src/app/order"
)

func (w *Wire) GetOrderCommandHandler() *order2.OrderCommandHandler {
	return order2.NewOrderCommandHandler(w.GetOrderRepository(), w.GetCoinRepository(), w.GetEventBus())
}

func (w *Wire) GetOrderRepository() *order2.OrderRepository {
	return order2.NewOrderRepository(w.GetRedisInfra())
}

func (w *Wire) GetOrderEventHandler() *order2.OrderEventHandler {
	return order2.NewOrderEventHandler(w.GetEventBus())
}
