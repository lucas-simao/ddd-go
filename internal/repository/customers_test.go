package repository

import (
	"context"
	"testing"

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
		FirstName: "Lucas",
		LastName:  "Simão",
		BirthDate: "1992-06-15",
	}
}

func (suite *CustomersTestSuite) TestPostCustomer() {
	tests := map[string]struct {
		Customer entity.Customer
		Error    error
	}{
		"Should register customer": {
			Customer: entity.Customer{
				FirstName: "lucas",
				LastName:  "simão",
				BirthDate: "1992-06-15",
			},
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			newCustomer, err := repo.PostCustomer(suite.ctx, test.Customer)
			suite.NoError(err)
			suite.Equal(test.Customer.FirstName, newCustomer.FirstName)
			suite.Equal(test.Customer.LastName, newCustomer.LastName)
			suite.Equal(test.Customer.BirthDate, newCustomer.BirthDate)
		})
	}
}
