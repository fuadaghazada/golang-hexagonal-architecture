package domain

import (
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/errs"
)

type Account struct {
	AccountId   string	`db:"account_id"`
	CustomerId  string	`db:"customer_id"`
	OpeningDate string	`db:"opening_date"`
	AccountType string	`db:"account_type"`
	Amount      float64	`db:"amount"`
	Status      string	`db:"status"`
}

func (a Account) ToResponseDto() dto.NewAccountResponseDto {
	return dto.NewAccountResponseDto{
		AccountId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return amount <= a.Amount
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(string) (*Account, *errs.AppError)
}