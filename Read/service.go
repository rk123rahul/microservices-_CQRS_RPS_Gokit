package main

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Customer struct {
	Customerid string `json:"customerid"`
	Email      string ` json:"email"`
	Phone      string ` json:"phone"`
}

type Repository interface {
	
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)

}

// service implements the ACcount Service
type accountservice struct {
	repository Repository
	logger     log.Logger
}

// Service describes the Account service.
type AccountService interface {
	
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	
}

// NewService creates and returns a new Account service instance
func NewService(rep Repository, logger log.Logger) AccountService {
	return &accountservice{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an customer

func (s accountservice) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetcustomerById")

	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetCustomerById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return customer, nil
}
func (s accountservice) GetAllCustomers(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllcustomers")
	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetAllCustomers(ctx)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return customer, nil
}

