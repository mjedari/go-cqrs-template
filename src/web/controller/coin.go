package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mjedari/go-cqrs-template/app/coin"
	"net/http"
	"strconv"
)

type CoinController struct {
	commandHandler coin.ICoinCommandHandler
	queryHandler   *coin.CoinQueryHandler
}

func NewCoinController(commandHandler coin.ICoinCommandHandler, queryHandler *coin.CoinQueryHandler) *CoinController {
	return &CoinController{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (c CoinController) CreateCoin(writer http.ResponseWriter, request *http.Request) {
	var command coin.CreateCoinCommand
	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	//if err := c.commandHandler.CreateCoin(context.Background(), command); err != nil {
	//	http.Error(writer, err.Error(), 500)
	//}

	c.commandHandler.Handle()

}

func (c CoinController) GetCoin(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	query := coin.GetCoinQuery{Id: uint(id)}
	coin, err := c.queryHandler.GetCoin(context.Background(), query)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	c.returnResponse(writer, coin)
}

func (c CoinController) GetCoinAll(writer http.ResponseWriter, request *http.Request) {

	query := coin.GetAllCoinsQuery{}
	users, err := c.queryHandler.GetCoins(context.Background(), query)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}

	c.returnResponse(writer, users)
}

func (c CoinController) returnResponse(writer http.ResponseWriter, response interface{}) {

	responseByte, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}

	_, err = writer.Write(responseByte)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}
}
