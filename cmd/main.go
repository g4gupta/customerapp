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
		Email: "rahul@mail.com",
	}
	customer2 := domain.Customer{
		ID:    "cust102",
		Name:  "Gaurav",
		Email: "gaurav@mail.com",
	}

	controller.Add(customer1)
	controller.Add(customer2)

	customer2update := domain.Customer{
		ID:    "cust102",
		Name:  "Gaurav",
		Email: "gaurav45@mail.com",
	}
	controller.Update(customer2update.ID, customer2update)
	controller.Get("cust102")
	controller.Get()
}
