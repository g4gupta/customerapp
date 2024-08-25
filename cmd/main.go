package main

import (
	"customerapp/domain"
	"customerapp/memstore"
)

func main() {
	controller := CustomerController{ // initialize customer controller
		repository: memstore.NewCustomerRepository(),
		//repository: mongodb.NewCustomerRepository(), // switching to another persistent store
	}
	customer1 := domain.Customer{
		ID:    "cust101",
		Name:  "Rahul",
		Email: "rahul@gmail.com",
	}
	controller.Add(customer1)
}
