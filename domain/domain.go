package domain

import (
	"errors"
	"fmt"
	"net/mail"
)

type Customer struct {
	ID    string
	Name  string
	Email string
}

var (
	ErrCustomerDoesNotExists = errors.New("the customer id is already present")
	ErrCustomerAlreadyExists = errors.New("the customer is not present")
	ErrKYCError              = errors.New("customer information is not correct")
)

type CustomerRepository interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() []Customer
}

// basic validations
func (cust *Customer) ValidateKYC() error {
	if _, err := mail.ParseAddress(cust.Email); err != nil {
		fmt.Println("Error: ", err)
		return ErrKYCError
	}
	return nil
}
