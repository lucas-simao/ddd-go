package api

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucas-simao/ddd-go/internal/domain/customers"
)

var port string = "9000"

func New(cs customers.Service) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	addRoutes(e, cs)

	return e
}

func Start(e *echo.Echo) {
	value, ok := os.LookupEnv("PORT")
	if ok {
		port = value
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
