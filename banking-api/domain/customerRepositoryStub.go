package domain

import "github.com/fuadaghazada/banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "101", Name: "Fuad", City: "Baku", Zipcode: "1010", DateOfBirth: "20/09/1998", Status: "ACTIVE"},
		{Id: "102", Name: "Ilyas", City: "Baku", Zipcode: "1012", DateOfBirth: "10/02/1998", Status: "SUSPENDED"},
	}

	return CustomerRepositoryStub{customers: customers}
}
