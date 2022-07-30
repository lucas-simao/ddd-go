package customers

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucas-simao/ddd-go/internal/entity"
)

type Service interface {
	PostCustomer(context.Context, entity.Customer) (entity.Customer, error)
	PutCustomer(context.Context, entity.Customer) (uuid.UUID, error)
	GetCustomerById(context.Context, uuid.UUID) (entity.Customer, error)
	DeleteCustomerById(context.Context, uuid.UUID) error
}
