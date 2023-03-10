package coin

import (
	"context"
	"github.com/mjedari/go-cqrs-template/domain/coin"
)

type CoinQueryHandler struct {
	repository *CoinRepository
}

func NewCoinQueryHandler(repository *CoinRepository) *CoinQueryHandler {
	return &CoinQueryHandler{repository: repository}
}

func (h CoinQueryHandler) GetCoin(ctx context.Context, query GetCoinQuery) (*coin.Coin, error) {
	coin := coin.Coin{Id: query.Id}

	if err := h.repository.GetCoin(ctx, &coin); err != nil {
		println("got err", err.Error())
		return nil, err
	}
	return &coin, nil
}

func (h CoinQueryHandler) GetCoins(ctx context.Context, query GetAllCoinsQuery) ([]coin.Coin, error) {

	coins, err := h.repository.GetAllCoins(ctx)
	if err != nil {
		return nil, err
	}

	return coins, nil
}
