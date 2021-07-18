package dto

import (
	"github.com/fuadaghazada/banking/errs"
	"strings"
)

type NewAccountDto struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountDto) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("Amount should be minimum 5000")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Amount type should be 'checking' or 'saving'")
	}

	return nil
}
