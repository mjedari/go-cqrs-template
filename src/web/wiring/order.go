package wiring

import "github.com/mjedari/go-cqrs-template/app/order"

func (w *Wire) GetOrderCommandHandler() *order.OrderCommandHandler {
	return order.NewOrderCommandHandler(w.GetOrderRepository(), w.GetCoinRepository(), w.GetEventBus())
}

func (w *Wire) GetOrderQueryHandler() *order.OrderQueryHandler {
	return order.NewOrderQueryHandler(w.GetOrderRepository())
}

func (w *Wire) GetOrderRepository() *order.OrderRepository {
	return order.NewOrderRepository(w.GetRedisInfra())
}

func (w *Wire) GetOrderEventHandler() *order.OrderEventHandler {
	return order.NewOrderEventHandler(w.GetEventBus())
}
