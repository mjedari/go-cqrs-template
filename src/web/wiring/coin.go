package wiring

import (
	"github.com/mjedari/go-cqrs-template/app/coin"
)

func (w *Wire) GetCoinCommandHandler() *coin.CoinCommandHandler {
	return coin.NewCoinCommandHandler(w.GetCoinRepository())
}

func (w *Wire) GetCoinRepository() *coin.CoinRepository {
	return coin.NewCoinRepository(w.GetRedisInfra())
}

func (w *Wire) GetCoinQueryHandler() *coin.CoinQueryHandler {
	return coin.NewCoinQueryHandler(w.GetCoinRepository())
}
