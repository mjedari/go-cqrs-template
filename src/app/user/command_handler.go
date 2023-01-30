package user

import (
	"context"
	userDomain "github.com/mjedari/go-cqrs-template/domain/user"
)

type UserCommandHandler struct {
	repository *UserRepository
}

func NewUserCommandHandler(repository *UserRepository) *UserCommandHandler {
	return &UserCommandHandler{repository: repository}
}

func (u UserCommandHandler) CreateUser(ctx context.Context, command CreateUserCommand) error {

	user := userDomain.NewUser(command.Name, command.Balance)

	if err := u.repository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
