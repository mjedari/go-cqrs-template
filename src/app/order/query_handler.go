package order

import (
	"context"
	"github.com/mjedari/go-cqrs-template/domain/coin"
)

type OrderQueryHandler struct {
	orderRepository *OrderRepository
}

func NewOrderQueryHandler(repository *OrderRepository) *OrderQueryHandler {
	return &OrderQueryHandler{orderRepository: repository}
}

func (h OrderQueryHandler) GetCoin(ctx context.Context, query OrderQueryHandler) (*coin.Coin, error) {

	return nil, nil
}
