package order

import (
	"context"
	"fmt"
	appCoin "github.com/mjedari/go-cqrs-template/app/coin"
	"github.com/mjedari/go-cqrs-template/domain/coin"
	"github.com/mjedari/go-cqrs-template/domain/order"
	"github.com/mjedari/go-cqrs-template/infra/providers/messaging"
)

type OrderCommandHandler struct {
	repository     *OrderRepository
	coinRepository *appCoin.CoinRepository
	eventBus       *messaging.EventBus
}

func NewOrderCommandHandler(repository *OrderRepository, coinRepository *appCoin.CoinRepository, eventBus *messaging.EventBus) *OrderCommandHandler {
	return &OrderCommandHandler{repository: repository, coinRepository: coinRepository, eventBus: eventBus}
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
	// get coin by coin Id
	if err := o.coinRepository.GetCoin(ctx, &coin); err != nil {
		return err
	}
	fmt.Println("your coin: ", coin)
	return nil
	// transaction initialized
	newOrder := order.NewOrder(command.CoinId, command.UserId, command.Amount)
	if err := o.repository.CreateOrder(ctx, newOrder); err != nil {
		return err
	}

	// settle if command is valid
	var initOrders order.InitializedOrders
	if err := o.repository.GetInitializedOrders(ctx, &initOrders); err != nil {
		//
	}

	// fire event
	// or
	if initOrders.GetTotalAmount() > coin.Min {
		//call settle function resuest to forign server
	}

	event := order.OrderEvent{
		CoinId:  1,
		OrderId: 1,
		UserId:  1,
	}
	if err := o.eventBus.Publish(ctx, []interface{}{event}); err != nil {

	}
	return nil
}

//func (o OrderCommandHandler) Settle(ctx context.Context, command SettleOrderCommand) error {
//
//}
