package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/ddd-go/internal/api/handlers"
	"github.com/lucas-simao/ddd-go/internal/domain/customers"
)

func addRoutes(e *echo.Echo, cs customers.Service) {
	v1 := e.Group("/v1")

	v1.POST("/customers", handlers.PostCustomer(cs))
	v1.GET("/customers/:id", handlers.GetCustomerById(cs))

}
