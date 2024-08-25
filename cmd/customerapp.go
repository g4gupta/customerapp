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
