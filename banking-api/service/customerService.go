package service

import (
	"github.com/fuadaghazada/banking/domain"
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerDto, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		return s.repo.FindAll("1")
	} else if status == "inactive" {
		return s.repo.FindAll("0")
	}
	return s.repo.FindAll("")
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerDto, *errs.AppError) {
	customer, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	customerDto := customer.ToDto()

	return &customerDto, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
