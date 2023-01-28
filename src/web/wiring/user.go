package wiring

import (
	"github.com/mjedari/go-cqrs-template/app/user"
)

func (w *Wire) GetUserCommandHandler() *user.UserCommandHandler {
	return user.NewUserCommandHandler(w.GetUserRepository())
}

func (w *Wire) GetUserQueryHandler() *user.UserQueryHandler {
	return user.NewUserQueryHandler(w.GetUserRepository())
}

func (w *Wire) GetUserRepository() *user.UserRepository {
	return user.NewUserRepository(w.GetRedisInfra())
}
