package order

import (
	"context"
	orderDomain "github.com/mjedari/go-cqrs-template/domain/order"
)

type OrderQueryHandler struct {
	repository *OrderRepository
}

func NewOrderQueryHandler(repository *OrderRepository) *OrderQueryHandler {
	return &OrderQueryHandler{repository: repository}
}

func (h OrderQueryHandler) GetOrder(ctx context.Context, query GetOrderQuery) (*orderDomain.Order, error) {
	order := orderDomain.Order{Id: query.Id}

	if err := h.repository.GetOrder(ctx, &order); err != nil {
		println("got err", err.Error())
		return nil, err
	}
	return &order, nil
}

func (u OrderQueryHandler) GetOrders(ctx context.Context, query GetAllOrdersQuery) ([]orderDomain.Order, error) {

	orders, err := u.repository.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
