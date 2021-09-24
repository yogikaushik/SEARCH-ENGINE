package routes

import (
	"github.com/yogeshkaushik1904/search-engine/controller"
	router "github.com/yogeshkaushik1904/search-engine/http"
)

var RegisterRoutes = func(httpRouter router.Router, c controller.OrderController) {

	httpRouter.GET("/orders/", c.GetOrders)
	httpRouter.GET("/orders/{id:[A-Za-z0-9_-]+}", c.GetOrderById)
	httpRouter.GET("/orders/search/", c.SearchOrders)

	httpRouter.POST("/orders/add/", c.AddOrder)
	httpRouter.PUT("/orders/update/{id:[A-Za-z0-9_-]+}", c.UpdateOrder)
	httpRouter.DELETE("/orders/delete/{id:[A-Za-z0-9_-]+}", c.DeleteOrder)
}
