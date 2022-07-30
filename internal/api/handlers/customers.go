package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/ddd-go/internal/domain/customers"
	"github.com/lucas-simao/ddd-go/internal/entity"
)

func PostCustomer(cs customers.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := entity.Customer{}

		ctx := c.Request().Context()

		err := c.Bind(&p)
		if err != nil {
			result.Message = "error to bind body: " + err.Error()
			return c.JSON(http.StatusBadRequest, err)
		}

		err = p.Validate()
		if err != nil {
			result.Message = "error to validate: " + err.Error()
			return c.JSON(http.StatusBadRequest, err)
		}

		newCustomer, err := cs.PostCustomer(ctx, p)
		if err != nil {
			result.Message = "error to create customer: " + err.Error()
			return c.JSON(http.StatusBadRequest, result)
		}

		return c.JSON(http.StatusCreated, newCustomer)
	}
}
