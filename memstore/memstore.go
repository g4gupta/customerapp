package memstore

import (
	"customerapp/domain"
	"errors"
)

var (
	ErrEmptyRepository = errors.New("No data in Customer Repository")
)

// CustomerRepository provides memory based local data repository.
type CustomerRepository struct {
	repository map[string]domain.Customer
}

// NewCustomerRepository initialises the memory data repository
func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{repository: make(map[string]domain.Customer)}
}

func (c *CustomerRepository) Create(customer domain.Customer) error {
	if _, ok := c.repository[customer.ID]; ok {
		return domain.ErrCustomerAlreadyExists
	} else {
		if err := customer.ValidateKYC(); err != nil {
			return err
		} else {
			c.repository[customer.ID] = customer
		}
	}
	return nil
}

func (c *CustomerRepository) Delete(custid string) error {
	if _, ok := c.repository[custid]; ok {
		delete(c.repository, custid)
	} else {
		return domain.ErrCustomerDoesNotExists
	}
	return nil
}

func (c *CustomerRepository) Update(custid string, customer domain.Customer) error {
	if _, ok := c.repository[custid]; ok {
		if err := customer.ValidateKYC(); err != nil {
			return err
		} else {
			c.repository[custid] = customer
		}
	} else {
		return domain.ErrCustomerDoesNotExists
	}
	return nil
}

func (c *CustomerRepository) GetById(custid string) (domain.Customer, error) {
	if customer, ok := c.repository[custid]; ok {
		return customer, nil
	} else {
		return domain.Customer{}, domain.ErrCustomerDoesNotExists

	}
}

func (c *CustomerRepository) GetAll() ([]domain.Customer, error) {
	//var resultset [len(c.repository)]domain.Customer
	if len(c.repository) == 0 {
		return nil, ErrEmptyRepository
	} else {
		resultset := make([]domain.Customer, 0)
		for _, v := range c.repository {
			resultset = append(resultset, v)
		}
		return resultset, nil
	}
}
