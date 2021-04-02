package main

import (
	"context"
	"database/sql"
	"errors"
	

	"github.com/go-kit/kit/log"
)

var (
	RepoErr             = errors.New("Unable to handle Repo Request")
	ErrIdNotFound       = errors.New("Id not found")
	ErrPhonenumNotFound = errors.New("Phone num is not found")
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mysql"),
	}, nil
}


func (repo *repo) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	customer := Customer{}

	err := repo.db.QueryRowContext(ctx, "SELECT c.customerid,c.email,c.phone FROM Customer as c where c.customerid = ?", id).Scan(&customer.Customerid, &customer.Email, &customer.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}
		return customer, err
	}
	return customer, nil
}
func (repo *repo) GetAllCustomers(ctx context.Context) (interface{}, error) {
	customer := Customer{}
	var res []interface{}
	rows, err := repo.db.QueryContext(ctx, "SELECT c.customerid,c.email,c.phone FROM Customer as c ")
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}
		return customer, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&customer.Customerid, &customer.Email, &customer.Phone)
		res = append([]interface{}{customer}, res...)
	}
	return res, nil
}

