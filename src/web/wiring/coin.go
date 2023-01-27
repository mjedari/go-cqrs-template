package wiring

import (
	"github.com/mjedari/go-cqrs-template/src/app/coin"
)

func (w *Wire) GetCoinRepository() *coin.CoinRepository {
	return coin.NewCoinRepository()
}
