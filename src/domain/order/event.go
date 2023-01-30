package order

type TestEvent struct {
	Name string
	Age  uint
}

type OrderEvent struct {
	CoinId   uint
	OrderId  uint
	UserId   uint
	Quantity uint
}

type InstantOrderEvent struct {
	CoinId   uint
	OrderId  uint
	UserId   uint
	Quantity uint
}

type FailTransactionEvent struct {
	Orders []Order
}
