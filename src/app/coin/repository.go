package coin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mjedari/go-cqrs-template/app/providers/storage"
	"github.com/mjedari/go-cqrs-template/domain/coin"
)

type IRepository interface {
	//
}

type CoinRepository struct {
	// pointer to gorm
	// pointer to redis
	redisStorage storage.IStorage
}

func NewCoinRepository(redisStorage storage.IStorage) *CoinRepository {
	return &CoinRepository{redisStorage: redisStorage}
}

func (r CoinRepository) GetCoin(ctx context.Context, coin *coin.Coin) error {
	key := coin.GetKey()
	coinByte, err := r.redisStorage.Select(ctx, key)
	if err != nil {
		return err
	}
	if mErr := json.Unmarshal(coinByte, coin); mErr != nil {
		return mErr
	}

	fmt.Println("this is coin", *coin)
	return nil
}

func (r CoinRepository) CreateCoin(ctx context.Context, coin *coin.Coin) error {
	key, err := r.attachNewId(ctx, coin)
	if err != nil {
		return err
	}
	fmt.Println("this is coin,", *coin)
	coinByte, err := json.Marshal(coin)
	if err != nil {
		return err
	}

	if err = r.redisStorage.Insert(ctx, key, coinByte); err != nil {
		return err
	}

	return nil
}

func (r CoinRepository) addId(ctx context.Context, coin *coin.Coin) error {
	coinId, err := r.redisStorage.GetNextId(ctx, coin)
	if err != nil {
		return err
	}

	coin.Id = uint(coinId)
	return nil

}

func (r CoinRepository) attachNewId(ctx context.Context, coin *coin.Coin) (string, error) {
	err := r.addId(ctx, coin)
	if err != nil {
		return "", err
	}
	key := coin.GetKey()
	return key, nil
}
