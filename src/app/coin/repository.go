package coin

import (
	"context"
	"github.com/mjedari/go-cqrs-template/src/domain/coin"
)

type IRepository interface {
	//
}

type CoinRepository struct {
	// pointer to gorm
	// pointer to redis
}

func NewCoinRepository() *CoinRepository {
	return &CoinRepository{}
}

func (r CoinRepository) GetCoin(ctx context.Context, coin *coin.Coin) error {
	coin.Min = 10
	coin.Name = "ABAN"

	return nil
}
