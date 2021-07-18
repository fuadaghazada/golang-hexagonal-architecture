package service

import (
	"github.com/fuadaghazada/banking/domain"
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountDto) (*dto.NewAccountResponseDto, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountDto) (*dto.NewAccountResponseDto, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	createdAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	accountResponseDto := createdAccount.ToResponseDto()

	return &accountResponseDto, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}