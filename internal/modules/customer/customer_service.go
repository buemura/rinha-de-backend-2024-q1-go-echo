package customer

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerService struct {
	db *pgxpool.Pool
}

func NewCustomerService(db *pgxpool.Pool) *CustomerService {
	return &CustomerService{
		db: db,
	}
} 

func (s *CustomerService) GetCustomer(customerID int) (*Customer, error) {
	rows, err := s.db.Query(
		context.Background(),
		"SELECT id, account_limit, account_balance FROM customers WHERE id = $1",
		customerID,
	)
	if err != nil {
		return nil, err
	}

	customerBalance, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[Customer])
	if err != nil {
		return nil, err
	}
	return customerBalance, nil
}
