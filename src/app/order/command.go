package order

type OrderCoinCommand struct {
	UserId uint
	CoinId uint
	Amount float64
}

type SettleOrderCommand struct {
	OrderId uint
}

type TestCommand struct {
	Name string
	Age  uint
}
