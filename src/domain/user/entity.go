package user

import "fmt"

type User struct {
	Id      uint    `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func NewUser(name string, balance float64) *User {
	return &User{Name: name, Balance: balance}
}

func (u *User) GetKey() string {
	return fmt.Sprintf("%v:%v", "user", u.Id)
}

func (u *User) HasCredit(amount float64) bool {
	return u.Balance > amount
}

func (u *User) SubtractBalance(amount float64) {
	u.Balance -= amount
}
