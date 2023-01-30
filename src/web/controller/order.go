package controller

import (
	"context"
	"encoding/json"
	"fmt"
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
	var command order.OrderCoinCommand

	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println("command", command)
	if err := o.commandHandler.OrderCoin(context.Background(), command); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("order has been received successfully"))
}

func (o OrderController) GetAllOrders(writer http.ResponseWriter, request *http.Request) {

	var query order.GetAllOrdersQuery
	orders, err := o.queryHandler.GetOrders(context.Background(), query)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}

	o.returnResponse(writer, orders)
}

func (o OrderController) returnResponse(writer http.ResponseWriter, response interface{}) {

	responseByte, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}

	_, err = writer.Write(responseByte)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}
}
