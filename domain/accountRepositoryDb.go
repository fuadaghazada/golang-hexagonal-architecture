package domain

import (
	"github.com/fuadaghazada/banking/errs"
	"github.com/fuadaghazada/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating a new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	accountId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the created account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(accountId, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}