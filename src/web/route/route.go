package route

import (
	"github.com/gorilla/mux"
	"github.com/mjedari/go-cqrs-template/web/controller"
	"github.com/mjedari/go-cqrs-template/web/middleware"
	"github.com/mjedari/go-cqrs-template/web/wiring"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	var userController = controller.NewUserController(wiring.Wiring.GetUserCommandHandler(), wiring.Wiring.GetUserQueryHandler())
	var orderController = controller.NewOrderController(wiring.Wiring.GetOrderCommandHandler(), wiring.Wiring.GetOrderQueryHandler())
	var coinController = controller.NewCoinController(wiring.Wiring.GetCoinCommandHandler(), wiring.Wiring.GetCoinQueryHandler())
	// user routes
	router.HandleFunc("/user/create", userController.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id:[0-9]+}", userController.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/all", userController.GetUserAll).Methods(http.MethodGet)
	// order routes
	router.HandleFunc("/order/create", orderController.OrderCoin).Methods(http.MethodPost)
	router.HandleFunc("/order/all", orderController.GetAllOrders).Methods(http.MethodGet)
	// coin routes
	router.HandleFunc("/coin/{id:[0-9]+}", coinController.GetCoin).Methods(http.MethodGet)
	router.HandleFunc("/coin/create", coinController.CreateCoin).Methods(http.MethodPost)
	router.HandleFunc("/coin/all", coinController.GetCoinAll).Methods(http.MethodGet)
	router.Use(middleware.LoggingMiddleware)

	return router
}
