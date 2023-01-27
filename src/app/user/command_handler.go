package user

import "context"

type UserCommandHandler struct {
	repository IRepository
}

func NewUserCommandHandler() *UserCommandHandler {
	//return &UserCommandHandler{repository: repository}
	return &UserCommandHandler{}
}

func (u UserCommandHandler) CreateUser(ctx context.Context, command CreateUserCommand) error {

	return nil
}
