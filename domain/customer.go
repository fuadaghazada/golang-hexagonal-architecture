package domain

import (
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}

func (customer Customer) ToDto() dto.CustomerDto {
	return dto.CustomerDto{
		Id:          customer.Id,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.statusAsText(),
	}
}

func (customer Customer) statusAsText() string {
	statusAsText := "active"
	if customer.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}
