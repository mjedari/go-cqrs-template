package order

import "time"

type OrderStatus string

const INITIATE OrderStatus = "INITIATE"
const SETTLED OrderStatus = "SETTLED"
const FAILED OrderStatus = "FAILED"

type Order struct {
	UserId    uint
	CoinId    uint
	Amount    float64
	Status    OrderStatus
	CreatedAt time.Time
}

type InitializedOrders []Order

func (orders InitializedOrders) GetTotalAmount() float64 {
	var amount float64
	for _, order := range orders {
		amount += order.Amount
	}
	return amount
}

func NewOrder(userId uint, coinId uint, amount float64) *Order {
	order := &Order{UserId: userId, CoinId: coinId, Amount: amount}
	order.Status = INITIATE
	order.CreatedAt = time.Now()
	return order
}

func (order Order) CanSettle() bool {
	return order.Amount > 10
}
