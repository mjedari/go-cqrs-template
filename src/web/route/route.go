package route

import (
	"github.com/gorilla/mux"
	"github.com/mjedari/go-cqrs-template/src/api/controller"
	"github.com/mjedari/go-cqrs-template/src/api/middleware"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	var userService = controller.NewUserService()
	//var coinService = service.NewCoinService()
	var orderService = controller.NewOrderService()
	//
	router.HandleFunc("/user/create", userService.CreateUser).Methods(http.MethodGet)
	router.HandleFunc("/order/coin", orderService.BuyCoin).Methods(http.MethodPost)
	router.Use(middleware.LoggingMiddleware)

	return router
}
