package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/lucas-simao/ddd-go/internal/entity"
	"github.com/stretchr/testify/suite"
)

type CustomersTestSuite struct {
	suite.Suite
	ctx      context.Context
	customer entity.Customer
}

func TestCustomersTestSuite(t *testing.T) {
	suite.Run(t, new(CustomersTestSuite))
}

func (suite *CustomersTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	suite.customer = entity.Customer{
		FirstName: "lucas",
		LastName:  "simão",
		BirthDate: "1992-06-15",
	}
}

func (suite *CustomersTestSuite) TestPostCustomer() {
	tests := map[string]struct {
		Customer entity.Customer
		Error    error
	}{
		"Should register customer": {
			Customer: suite.customer,
		},
		"Should return error": {
			Customer: entity.Customer{},
			Error:    entity.PostCustomerError,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			newCustomer, err := repo.PostCustomer(suite.ctx, test.Customer)
			if err != nil {
				suite.Equal(test.Error, err)
				return
			}
			suite.NoError(err)
			suite.NotEmpty(newCustomer.Id.String())
			suite.Equal(test.Customer.FirstName, newCustomer.FirstName)
			suite.Equal(test.Customer.LastName, newCustomer.LastName)
			suite.Equal(test.Customer.BirthDate, newCustomer.BirthDate)
		})
	}
}

func (suite *CustomersTestSuite) TestGetCustomer() {
	newCustomer, err := repo.PostCustomer(suite.ctx, suite.customer)
	if err != nil {
		suite.Error(err)
	}

	tests := map[string]struct {
		Customer entity.Customer
		Error    error
	}{
		"Should get customer by id": {
			Customer: newCustomer,
			Error:    nil,
		},
		"Should get error wrong id": {
			Customer: entity.Customer{
				Id: uuid.New(),
			}, Error: entity.GetCustomerError,
		},
		"Should get error empty id": {
			Customer: entity.Customer{
				Id: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			}, Error: entity.GetCustomerError,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			result, err := repo.GetCustomerById(suite.ctx, test.Customer.Id)
			if err != nil {
				suite.Equal(test.Error, err)
				return
			}

			suite.NoError(err)
			suite.Equal(test.Customer.Id, result.Id)
			suite.Equal(test.Customer.FirstName, result.FirstName)
			suite.Equal(test.Customer.LastName, result.LastName)
			suite.Equal(test.Customer.BirthDate, result.BirthDate)
		})
	}
}

func (suite *CustomersTestSuite) TestPutCustomer() {
	customer, err := repo.PostCustomer(suite.ctx, suite.customer)
	if err != nil {
		suite.Error(err)
	}

	tests := map[string]struct {
		newCustomerData entity.Customer
		Error           error
	}{
		"Should update customer": {
			newCustomerData: entity.Customer{
				Id:        customer.Id,
				FirstName: "joão",
				LastName:  "simão",
				BirthDate: "2019-11-14",
			},
			Error: nil,
		},
		"Should get error wrong id": {
			newCustomerData: entity.Customer{
				Id: uuid.New(),
			},
			Error: entity.PutCustomerError,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			result, err := repo.PutCustomerById(suite.ctx, test.newCustomerData)
			if err != nil {
				suite.Equal(test.Error, err)
				return
			}

			suite.NoError(err)
			suite.Equal(test.newCustomerData.Id, result.Id)
			suite.Equal(test.newCustomerData.FirstName, result.FirstName)
			suite.Equal(test.newCustomerData.LastName, result.LastName)
			suite.Equal(test.newCustomerData.BirthDate, result.BirthDate)
		})
	}
}

func (suite *CustomersTestSuite) TestDeleteCustomer() {
	customer, err := repo.PostCustomer(suite.ctx, suite.customer)
	if err != nil {
		suite.Error(err)
	}

	tests := map[string]struct {
		customer entity.Customer
		Error    error
	}{
		"Should update customer": {
			customer: customer,
			Error:    nil,
		},
		"Should get error wrong id": {
			customer: entity.Customer{
				Id: uuid.New(),
			},
			Error: entity.DeleteCustomerError,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			err := repo.DeleteCustomerById(suite.ctx, test.customer.Id)
			if err != nil {
				suite.Equal(test.Error, err)
				return
			}

			suite.NoError(err)
		})
	}
}
