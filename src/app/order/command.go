package order

import "github.com/mjedari/go-cqrs-template/domain/coin"

type OrderCoinCommand struct {
	UserId   uint `json:"user_id"`
	CoinId   uint `json:"coin_id"`
	Quantity uint `json:"quantity"`
}

type SettleOrderCommand struct {
	OrderId uint
}

type TestCommand struct {
	Name string
	Age  uint
}

type ExchangeTransaction struct {
	Coin     coin.Coin
	Quantity uint
}
