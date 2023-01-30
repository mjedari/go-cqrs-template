package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mjedari/go-cqrs-template/app/providers/storage"
	orderDomain "github.com/mjedari/go-cqrs-template/domain/order"
)

type IRepository interface {
	//
}

type OrderRepository struct {
	// pointer to gorm
	// pointer to redis
	redisStorage storage.IStorage
}

func NewOrderRepository(storage storage.IStorage) *OrderRepository {
	return &OrderRepository{storage}
}

func (o OrderRepository) CreateOrder(ctx context.Context, order *orderDomain.Order) error {
	key, err := o.attachNewId(ctx, order)
	if err != nil {
		return err
	}
	coinByte, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err = o.redisStorage.Insert(ctx, key, coinByte); err != nil {
		return err
	}
	return nil
}

func (o OrderRepository) GetOrder(ctx context.Context, order *orderDomain.Order) error {
	key := order.GetKey()

	orderByte, err := o.redisStorage.Select(ctx, key)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(orderByte, order); err != nil {
		return err
	}

	return nil
}

func (o OrderRepository) GetAllOrders(ctx context.Context) ([]orderDomain.Order, error) {
	allOrdersBytes, err := o.redisStorage.SelectAll(ctx, "order:*")
	if err != nil {
		return nil, err
	}

	var allOrders []orderDomain.Order
	for _, ordersByte := range allOrdersBytes {
		var order orderDomain.Order
		if errM := json.Unmarshal(ordersByte, &order); errM != nil {
			return nil, errM
		}
		allOrders = append(allOrders, order)
	}

	return allOrders, nil
}

func (o OrderRepository) UpdateOrder(ctx context.Context, order *orderDomain.Order) error {
	// store order here
	fmt.Println("Order updated in redis", order)
	return nil
}

func (o OrderRepository) SettleOrders(ctx context.Context, orders []orderDomain.Order) error {

	var updatedOrders []orderDomain.Order
	for _, order := range orders {
		order.Status = orderDomain.SETTLED
		updatedOrders = append(updatedOrders, order)
	}

	for _, order := range updatedOrders {
		key := order.GetKey()
		orderByte, err := json.Marshal(order)
		if err != nil {
			return err
		}

		if err := o.redisStorage.Update(ctx, key, orderByte); err != nil {
			return err
		}
	}

	return nil

}

func (o OrderRepository) GetInitializedOrders(ctx context.Context) (*orderDomain.InitializedOrders, error) {
	allOrdersBytes, err := o.redisStorage.SelectAll(ctx, "order:*")
	if err != nil {
		return nil, err
	}

	var allOrders []orderDomain.Order
	for _, ordersByte := range allOrdersBytes {
		var order orderDomain.Order
		if errM := json.Unmarshal(ordersByte, &order); errM != nil {
			return nil, errM
		}
		allOrders = append(allOrders, order)
	}

	var initOrders orderDomain.InitializedOrders
	for _, order := range allOrders {
		if order.Status == orderDomain.INITIATE {
			initOrders.List = append(initOrders.List, order)
		}
	}

	return &initOrders, nil
}

func (o OrderRepository) PingRedis() error {

	//ping := o.redis.Ping(context.Background())

	//fmt.Println(ping)
	return nil
}

func (o OrderRepository) addId(ctx context.Context, order *orderDomain.Order) error {
	orderId, err := o.redisStorage.GetNextId(ctx, order)
	if err != nil {
		return err
	}

	order.Id = uint(orderId)
	return nil

}

func (o OrderRepository) attachNewId(ctx context.Context, order *orderDomain.Order) (string, error) {
	err := o.addId(ctx, order)
	if err != nil {
		return "", err
	}
	key := order.GetKey()
	return key, nil
}
