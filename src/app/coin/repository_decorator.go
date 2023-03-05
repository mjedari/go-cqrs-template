package coin

import "fmt"

type RepositoryDecorator struct {
	handler *CoinCommandHandler
}

func (r RepositoryDecorator) Handle() {
	fmt.Println("RepositoryDecorator")
	r.handler.Handle()
}

func NewRepositoryDecorator(handler *CoinCommandHandler) *RepositoryDecorator {
	return &RepositoryDecorator{handler: handler}
}
