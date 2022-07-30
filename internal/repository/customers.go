package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucas-simao/ddd-go/internal/entity"
)

func (r *repository) PostCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	namedStmt, err := r.db.PrepareNamedContext(ctx, sqlCreateCustomer)
	if err != nil {
		return entity.Customer{}, err
	}

	var newCustomer entity.Customer

	err = namedStmt.GetContext(ctx, &newCustomer, customer)
	if err != nil {
		return entity.Customer{}, err
	}

	return newCustomer, nil
}

func (r *repository) PutCustomer(context.Context, entity.Customer) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (r *repository) GetCustomerById(context.Context, uuid.UUID) (entity.Customer, error) {
	return entity.Customer{}, nil
}

func (r *repository) DeleteCustomerById(context.Context, uuid.UUID) error {
	return nil
}
