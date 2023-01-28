package user

import "fmt"

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func (u User) GetKey() string {
	return fmt.Sprintf("%v:%v", "user", u.Id)
}
