package coin

import "fmt"

type Coin struct {
	Id    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Min   float64 `json:"min"`
}

func NewCoin(name string, price float64, min float64) *Coin {
	return &Coin{Name: name, Price: price, Min: min}
}

func (c Coin) GetKey() string {
	return fmt.Sprintf("%v:%v", "coin", c.Id)
}

func (c Coin) GetMinPrice() float64 {
	return c.Min * c.Price
}
