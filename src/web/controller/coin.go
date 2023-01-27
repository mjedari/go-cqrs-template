package controller

import (
	"github.com/mjedari/go-cqrs-template/src/app/coin"
)

type CoinService struct {
	commandHandler *coin.CoinCommandHandler
}

func NewCoinService() *CoinService {
	handler := coin.NewCoinCommandHandler()
	return &CoinService{commandHandler: handler}
}
