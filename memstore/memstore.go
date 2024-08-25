package memstore

import (
	"customerapp/domain"
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
	c.repository[customer.ID] = customer
	return nil
}

func (c *CustomerRepository) Delete(custid string) error {
	delete(c.repository, custid)
	return nil
}

func (c *CustomerRepository) Update(custid string, customer domain.Customer) error {
	c.repository[custid] = customer
	return nil
}

func (c *CustomerRepository) GetById(custid string) (domain.Customer, error) {
	return c.repository[custid], nil
}

func (c *CustomerRepository) GetAll() ([]domain.Customer, error) {
	resultset := make([]domain.Customer, len(c.repository))

	for _, v := range c.repository {
		resultset = append(resultset, v)
	}
	return resultset, nil
}
