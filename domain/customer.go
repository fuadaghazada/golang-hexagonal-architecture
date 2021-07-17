package domain

import "github.com/fuadaghazada/banking/errs"

type Customer struct {
	Id          string	`db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string	`db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
