package service

import "github.com/fuadaghazada/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, error) {
	return s.repo.FindById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
