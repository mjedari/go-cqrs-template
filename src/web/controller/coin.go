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
	commandHandler *coin.CoinCommandHandler
	queryHandler   *coin.CoinQueryHandler
}

func NewCoinController(commandHandler *coin.CoinCommandHandler, queryHandler *coin.CoinQueryHandler) *CoinController {
	return &CoinController{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (c CoinController) CreateCoin(writer http.ResponseWriter, request *http.Request) {
	var command coin.CreateCoinCommand
	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	if err := c.commandHandler.CreateCoin(context.Background(), command); err != nil {
		http.Error(writer, err.Error(), 500)
	}

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

	response, err := json.Marshal(coin)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	_, err = writer.Write(response)
	if err != nil {
		http.Error(writer, "internal error", 500)
	}
}
