package user

import (
	"context"
	userDomain "github.com/mjedari/go-cqrs-template/domain/user"
)

type UserQueryHandler struct {
	repository *UserRepository
}

func NewUserQueryHandler(repository *UserRepository) *UserQueryHandler {
	return &UserQueryHandler{repository: repository}
}

func (u UserQueryHandler) GetUser(ctx context.Context, command GetUserQuery) (*userDomain.User, error) {
	user := userDomain.User{
		Id: command.Id,
	}

	if err := u.repository.GetUser(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
