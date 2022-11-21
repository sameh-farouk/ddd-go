package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sameh-farouk/ddd-go/entity"
	"github.com/sameh-farouk/ddd-go/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid name")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	// Create a new person and generate ID
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	customer := Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}
	return customer, nil
}
