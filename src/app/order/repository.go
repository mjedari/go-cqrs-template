package order

import (
	"context"
	"fmt"
	"github.com/mjedari/go-cqrs-template/src/app/providers/storage"
	"github.com/mjedari/go-cqrs-template/src/domain/order"
)

type IRepository interface {
	//
}

type OrderRepository struct {
	// pointer to gorm
	// pointer to redis
	redisStorage storage.IRedisStorage
}

func NewOrderRepository(storage storage.IRedisStorage) *OrderRepository {
	return &OrderRepository{storage}
}

func (o OrderRepository) CreateOrder(ctx context.Context, order *order.Order) error {
	// store order here

	// insertDB
	fmt.Println("Order stored in redis", order)
	if err := o.redisStorage.Insert(context.Background(), "my-key", "my-value"); err != nil {
		return err
	}
	return nil
}

func (o OrderRepository) UpdateOrder(ctx context.Context, order *order.Order) error {
	// store order here
	fmt.Println("Order updated in redis", order)
	return nil
}

func (o OrderRepository) GetInitializedOrders(ctx context.Context, orders *order.InitializedOrders) error {
	// get all initilized orders
	return nil
}

func (o OrderRepository) PingRedis() error {

	//ping := o.redis.Ping(context.Background())

	//fmt.Println(ping)
	return nil
}
