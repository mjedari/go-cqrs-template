package controller

import (
	"context"
	"github.com/mjedari/go-cqrs-template/app/order"
	"net/http"
)

type OrderController struct {
	commandHandler *order.OrderCommandHandler
	queryHandler   *order.OrderQueryHandler
}

func NewOrderController(commandHandler *order.OrderCommandHandler, queryHandler *order.OrderQueryHandler) *OrderController {
	return &OrderController{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (o OrderController) OrderCoin(w http.ResponseWriter, request *http.Request) {
	command := order.OrderCoinCommand{
		UserId: 1,
		CoinId: 1,
		Amount: 10,
	}

	if err := o.commandHandler.OrderCoin(context.Background(), command); err != nil {
		http.Error(w, err.Error(), 500)
	}

	//var command order.TestCommand
	//if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
	//
	//}
	//
	//if err := o.commandHandler.OrderTest(context.Background(), command); err != nil {
	//	http.Error(w, err.Error(), 500)
	//}

}
