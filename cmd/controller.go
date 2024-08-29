package main

import (
	"fmt"

	"customerapp/domain"
)

// customnerapp user interface
type CustomerController struct {
	repository domain.CustomerRepository
}

// user interface fuction to create customer
func (cc CustomerController) Add(c domain.Customer) {
	err := cc.repository.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created")
}

// user interface fuction to update customer
func (cc CustomerController) Update(id string, c domain.Customer) {
	err := cc.repository.Update(id, c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer Id: ", id, "has been updated")
}

func (cc CustomerController) Get(custid ...string) {
	if custid == nil {
		if customers, err := cc.repository.GetAll(); err != nil {
			for _, cust := range customers {
				fmt.Println(cust.ID, "|", cust.Name, "|", cust.Email)
			}
		} else {
			fmt.Println("Error: ", err)
		}
	} else {
		for _, id := range custid {
			if customer, err := cc.repository.GetById(id); err == nil {
				fmt.Println(id, "|", customer.Name, "|", customer.Email)
			} else {
				fmt.Println("Error:", err)
			}
		}
	}
}
