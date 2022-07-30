package customers

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucas-simao/ddd-go/internal/entity"
	"github.com/lucas-simao/ddd-go/internal/repository"
)

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return service{
		repository: r,
	}
}

func (s service) PostCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	return s.repository.PostCustomer(ctx, customer)
}

func (s service) PutCustomer(ctx context.Context, customer entity.Customer) (customerId uuid.UUID, err error) {
	return
}

func (s service) GetCustomerById(ctx context.Context, customerId uuid.UUID) (customer entity.Customer, err error) {
	return s.repository.GetCustomerById(ctx, customerId)
}

func (s service) DeleteCustomerById(ctx context.Context, customerId uuid.UUID) (err error) {
	return
}
