package domain

import (
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToResponseDto() dto.NewAccountResponseDto {
	return dto.NewAccountResponseDto{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}