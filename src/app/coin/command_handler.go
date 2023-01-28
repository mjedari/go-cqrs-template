package coin

import (
	"context"
	coinDomain "github.com/mjedari/go-cqrs-template/domain/coin"
)

type CoinCommandHandler struct {
	repository *CoinRepository
}

func NewCoinCommandHandler(repository *CoinRepository) *CoinCommandHandler {
	return &CoinCommandHandler{repository: repository}
}

func (h CoinCommandHandler) CreateCoin(ctx context.Context, command CreateCoinCommand) error {

	coin := coinDomain.NewCoin(command.Name, command.Price, command.Min)
	if err := h.repository.CreateCoin(ctx, coin); err != nil {
		return err
	}

	return nil
}
