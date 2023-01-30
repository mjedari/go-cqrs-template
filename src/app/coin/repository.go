package coin

import (
	"context"
	"encoding/json"
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

	return nil
}

func (r CoinRepository) GetAllCoins(ctx context.Context) ([]coin.Coin, error) {
	allCoinsBytes, err := r.redisStorage.SelectAll(ctx, "coin:*")
	if err != nil {
		return nil, err
	}

	var allCoins []coin.Coin
	for _, coinByte := range allCoinsBytes {
		var coin coin.Coin
		if errM := json.Unmarshal(coinByte, &coin); errM != nil {
			return nil, errM
		}
		allCoins = append(allCoins, coin)
	}

	return allCoins, nil
}

func (r CoinRepository) CreateCoin(ctx context.Context, coin *coin.Coin) error {
	key, err := r.attachNewId(ctx, coin)
	if err != nil {
		return err
	}
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
