package order

import (
	"context"
	"errors"
	"fmt"
	appCoin "github.com/mjedari/go-cqrs-template/app/coin"
	"github.com/mjedari/go-cqrs-template/app/user"
	"github.com/mjedari/go-cqrs-template/domain/coin"
	"github.com/mjedari/go-cqrs-template/domain/order"
	userDomain "github.com/mjedari/go-cqrs-template/domain/user"
	"github.com/mjedari/go-cqrs-template/infra/providers/messaging"
	"github.com/sirupsen/logrus"
	"time"
)

type OrderCommandHandler struct {
	repository     *OrderRepository
	coinRepository *appCoin.CoinRepository
	userRepository *user.UserRepository
	eventBus       *messaging.EventBus
}

func NewOrderCommandHandler(repository *OrderRepository, coinRepository *appCoin.CoinRepository, userRepository *user.UserRepository, eventBus *messaging.EventBus) *OrderCommandHandler {
	return &OrderCommandHandler{repository: repository, coinRepository: coinRepository, userRepository: userRepository, eventBus: eventBus}
}

func (o OrderCommandHandler) OrderTest(ctx context.Context, command TestCommand) error {
	//dispatch and event
	event := order.TestEvent{
		Name: command.Name,
		Age:  command.Age,
	}
	if err := o.eventBus.Publish(ctx, []interface{}{event}); err != nil {
		return err
	}

	return nil
}

func (o OrderCommandHandler) OrderCoin(ctx context.Context, command OrderCoinCommand) error {
	coin := coin.Coin{Id: command.CoinId}

	if err := o.coinRepository.GetCoin(ctx, &coin); err != nil {
		return err
	}

	user := userDomain.User{Id: command.UserId}
	if err := o.userRepository.GetUser(ctx, &user); err != nil {
		return err
	}

	orderTotalPrice := float64(command.Quantity) * coin.Price
	if !user.HasCredit(orderTotalPrice) {
		return errors.New("user has not enough balance")
	}

	newOrder := order.NewOrder(order.OrderParams{
		UserId:    user.Id,
		CoinId:    coin.Id,
		Quantity:  command.Quantity,
		Amount:    orderTotalPrice,
		Status:    order.INITIATE,
		CreatedAt: time.Now(),
	})

	// todo: should handle transactional
	user.SubtractBalance(orderTotalPrice)
	if err := o.userRepository.UpdateUser(ctx, &user); err != nil {
		return err
	}

	if err := o.repository.CreateOrder(ctx, newOrder); err != nil {
		return err
	}

	var events []interface{}
	// todo: here

	initOrders, err := o.repository.GetInitializedOrders(ctx)
	if err != nil {
		// todo: handle err
		logrus.Error(err)
	}

	if initOrders.GetTotalAmount() > order.MinExchangeOrderAmount {
		fmt.Println("Got event fire", newOrder.Amount, newOrder.Amount)
		event := order.InstantOrderEvent{
			CoinId:  newOrder.CoinId,
			OrderId: newOrder.Id,
			UserId:  newOrder.UserId,
		}
		events = append(events, event)

		if err := o.eventBus.Publish(ctx, events); err != nil {
			return err
		}
	}

	return nil
}
