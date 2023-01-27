package order

type TestEvent struct {
	Name string
	Age  uint
}

type OrderEvent struct {
	CoinId  uint
	OrderId uint
	UserId  uint
}
