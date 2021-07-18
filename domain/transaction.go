package domain

import (
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/errs"
	"strings"
)

type Transaction struct {
	TransactionId   string	`db:"transaction_id"`
	AccountId       string	`db:"account_id"`
	Amount          float64	`db:"amount"`
	TransactionType string	`db:"transaction_type"`
	TransactionDate string	`db:"transaction_date"`
}

func (t Transaction) ToResponseDto(balance float64) dto.NewTransactionResponseDto {
	return dto.NewTransactionResponseDto{
		TransactionId: t.TransactionId,
		AccountId: t.AccountId,
		Balance: balance,
	}
}

func (t Transaction) IsWithdraw() bool {
	return strings.ToLower(t.TransactionType) == dto.WITHDRAW
}

type TransactionRepository interface {
	Save(transaction Transaction) (*Transaction, *errs.AppError)
}