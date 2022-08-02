package entity

import (
	"errors"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/google/uuid"
)

var (
	PostCustomerError   = errors.New("Error to register customer")
	GetCustomerError    = errors.New("Error to get customer")
	PutCustomerError    = errors.New("Error to update customer")
	DeleteCustomerError = errors.New("Error to delete customer")
)

type Customer struct {
	Id        uuid.UUID  `json:"id" db:"id"`
	FirstName string     `json:"firstName" db:"first_name"`
	LastName  string     `json:"lastName" db:"last_name"`
	BirthDate string     `json:"birthDate" db:"birth_date"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

func (c Customer) Validate() error {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.FirstName,
			validation.Required,
			validation.Length(1, 50),
			validation.By(func(name interface{}) error {
				s, _ := name.(string)
				if len(strings.Split(s, " ")) > 1 {
					return errors.New("the field firstName only support one name")
				}
				return nil
			})),
		validation.Field(&c.LastName,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(&c.BirthDate,
			validation.Required, validation.Date("2006-01-02"),
		))

	if err != nil {
		return err
	}

	return nil
}

// type Address struct {
// 	Id           uuid.UUID
// 	Name         string
// 	Street       string
// 	Neighborhood string
// 	City         string
// 	State        string
// 	Country      string
// 	PostalCode   string
// }

// type Contacts struct {
// 	Phone int
// 	Email string
// }
