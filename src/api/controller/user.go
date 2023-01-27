package controller

import (
	"github.com/mjedari/go-cqrs-template/src/app/user"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserService struct {
	commandHandler *user.UserCommandHandler
}

func NewUserService() *UserService {
	handler := user.NewUserCommandHandler()
	return &UserService{commandHandler: handler}
}

func (u UserService) CreateUser(w http.ResponseWriter, request *http.Request) {
	logrus.Info("Create user api called: ")
}
