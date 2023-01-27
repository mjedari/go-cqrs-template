package controller

import (
	"context"
	"encoding/json"
	"github.com/mjedari/go-cqrs-template/src/api/wiring"
	order2 "github.com/mjedari/go-cqrs-template/src/app/order"
	"github.com/sirupsen/logrus"
	"net/http"
)

type OrderService struct {
	commandHandler *order2.OrderCommandHandler
}

func NewOrderService() *OrderService {
	handler := wiring.Wiring.GetOrderCommandHandler()
	return &OrderService{commandHandler: handler}
}

func (o OrderService) BuyCoin(w http.ResponseWriter, request *http.Request) {
	logrus.Info("BuyCoin api called:")

	//command := order.OrderCoinCommand{
	//	UserId: 1,
	//	CoinId: 1,
	//	Amount: 10,
	//}
	//
	//if err := o.commandHandler.OrderCoin(context.Background(), command); err != nil {
	//	http.Error(w, err.Error(), 500)
	//}

	var command order2.TestCommand
	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {

	}

	if err := o.commandHandler.OrderTest(context.Background(), command); err != nil {
		http.Error(w, err.Error(), 500)
	}

}
