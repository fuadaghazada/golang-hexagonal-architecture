package domain

import (
	"database/sql"
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

func (d AccountRepositoryDb) FindById(accountId string) (*Account, *errs.AppError) {
	sqlGet := "SELECT customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?"

	var account Account
	err := d.client.Get(&account, sqlGet, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		}

		logger.Error("Error while getting the account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}