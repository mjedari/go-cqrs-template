package order

import "github.com/mjedari/go-cqrs-template/src/domain/coin"

type IOrderEvent interface {
	GetName() string
	GetAmount() float64
	GetCoin() coin.Coin
}
