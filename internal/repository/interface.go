package repository

import (
	"context"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/lucas-simao/ddd-go/internal/entity"
)

type Repository interface {
	PostCustomer(context.Context, entity.Customer) (entity.Customer, error)
	PutCustomer(context.Context, entity.Customer) (uuid.UUID, error)
	GetCustomerById(context.Context, uuid.UUID) (entity.Customer, error)
	DeleteCustomerById(context.Context, uuid.UUID) error
}
