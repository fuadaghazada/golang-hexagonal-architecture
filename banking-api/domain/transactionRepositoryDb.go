package domain

import (
	"github.com/fuadaghazada/banking/errs"
	"github.com/fuadaghazada/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDb) Save(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while creating a new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, _ := tx.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdraw() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			logger.Error("Error while rollback: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected error from database")
		}
		logger.Error("Error while saving the new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the created account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	err = tx.Commit()
	if err != nil {
		logger.Error("Error while saving the new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &t, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{client: dbClient}
}
