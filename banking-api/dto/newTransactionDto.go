package dto

import (
	"github.com/fuadaghazada/banking/errs"
	"strings"
)

const WITHDRAW = "withdraw"
const DEPOSIT = "deposit"

type NewTransactionDto struct {
	AccountId       string
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (dto NewTransactionDto) Validate() *errs.AppError {
	if strings.ToLower(dto.TransactionType) != WITHDRAW && strings.ToLower(dto.TransactionType) != DEPOSIT {
		return errs.NewValidationError("Invalid transaction type")
	}
	if dto.Amount < 0 {
		return errs.NewValidationError("Transaction amount cannot be negative")
	}
	return nil
}

func (dto NewTransactionDto) IsTypeWithdrawal() bool {
	return strings.ToLower(dto.TransactionType) == WITHDRAW
}
