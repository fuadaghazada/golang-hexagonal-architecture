package domain

import (
	"database/sql"
	"github.com/fuadaghazada/banking/errs"
	"github.com/fuadaghazada/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying the customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}

		logger.Error("Error while scanning the customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: dbClient}
}