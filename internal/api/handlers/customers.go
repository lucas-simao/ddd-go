package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/ddd-go/internal/domain/customers"
	"github.com/lucas-simao/ddd-go/internal/entity"
)

func PostCustomer(cs customers.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		p := entity.Customer{}

		err := c.Bind(&p)
		if err != nil {
			result.Message = fmt.Sprintf("error to bind body: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		err = p.Validate()
		if err != nil {
			result.Message = fmt.Sprintf("error to validate: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		newCustomer, err := cs.PostCustomer(ctx, p)
		if err != nil {
			result.Message = fmt.Sprintf("error to create customer: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		return c.JSON(http.StatusCreated, newCustomer)
	}
}

func GetCustomerById(cs customers.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		customerId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			result.Message = fmt.Sprintf("error to get customer id: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		customer, err := cs.GetCustomerById(ctx, customerId)
		if err != nil {
			result.Message = fmt.Sprintf("error to get customer: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		return c.JSON(http.StatusOK, customer)
	}
}

func PutCustomerById(cs customers.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		p := entity.Customer{}

		err := c.Bind(&p)
		if err != nil {
			result.Message = fmt.Sprintf("error to bind body: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		err = p.Validate()
		if err != nil {
			result.Message = fmt.Sprintf("error to validate: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		customerId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			result.Message = fmt.Sprintf("error to get customer id: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		customer, err := cs.GetCustomerById(ctx, customerId)
		if err != nil {
			result.Message = fmt.Sprintf("error to get customer: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		if customer.Id == uuid.Nil {
			result.Message = "customer not exist"
			return c.JSON(http.StatusBadRequest, result)
		}

		p.Id = customerId

		customerUpdated, err := cs.PutCustomerById(ctx, p)
		if err != nil {
			result.Message = fmt.Sprintf("error to update customer: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		return c.JSON(http.StatusOK, customerUpdated)
	}
}

func DeleteCustomerById(cs customers.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		customerId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			result.Message = fmt.Sprintf("error to get customer id: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		customer, err := cs.GetCustomerById(ctx, customerId)
		if err != nil {
			result.Message = fmt.Sprintf("error to get customer: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		if customer.Id == uuid.Nil {
			result.Message = "customer not exist"
			return c.JSON(http.StatusBadRequest, result)
		}

		err = cs.DeleteCustomerById(ctx, customerId)
		if err != nil {
			result.Message = fmt.Sprintf("error to delete customer: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		return c.NoContent(http.StatusNoContent)
	}
}
