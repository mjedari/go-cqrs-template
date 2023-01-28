package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mjedari/go-cqrs-template/app/user"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserController struct {
	commandHandler *user.UserCommandHandler
	queryHandler   *user.UserQueryHandler
}

func NewUserController(commandHandler *user.UserCommandHandler, queryHandler *user.UserQueryHandler) *UserController {
	return &UserController{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (u UserController) CreateUser(w http.ResponseWriter, request *http.Request) {
	logrus.Info("Create user web called: ")

	var command user.CreateUserCommand
	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := u.commandHandler.CreateUser(context.Background(), command); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (u UserController) GetUser(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	query := user.GetUserQuery{Id: uint(id)}
	user, err := u.queryHandler.GetUser(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if user == nil {
		http.NotFound(w, request)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
