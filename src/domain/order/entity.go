package order

import (
	"fmt"
	"time"
)

type OrderStatus string

const INITIATE OrderStatus = "INITIATE"
const SETTLED OrderStatus = "SETTLED"
const FAILED OrderStatus = "FAILED"

const MinExchangeOrderAmount float64 = 10

type Order struct {
	Id        uint        `json:"id"`
	UserId    uint        `json:"user_id"`
	CoinId    uint        `json:"coin_id"`
	Quantity  uint        `json:"quantity"`
	Amount    float64     `json:"amount"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

func NewOrder(params OrderParams) *Order {
	return &Order{
		UserId:    params.UserId,
		CoinId:    params.CoinId,
		Quantity:  params.Quantity,
		Amount:    params.Amount,
		Status:    params.Status,
		CreatedAt: params.CreatedAt,
	}
}

func (order *Order) GetKey() string {
	return fmt.Sprintf("%v:%v", "order", order.Id)
}

type OrderParams struct {
	UserId    uint
	CoinId    uint
	Quantity  uint
	Amount    float64
	Status    OrderStatus
	CreatedAt time.Time
}

type InitializedOrders struct {
	List []Order
}

func (orders InitializedOrders) GetTotalAmount() float64 {
	var amount float64
	for _, order := range orders.List {
		amount += order.Amount
	}
	return amount
}

func (orders InitializedOrders) GetTotalQuantity() uint {
	var quantity uint
	for _, order := range orders.List {
		quantity += order.Quantity
	}
	return quantity
}

func (order *Order) ChangeStatus(status OrderStatus) {
	order.Status = status
}
