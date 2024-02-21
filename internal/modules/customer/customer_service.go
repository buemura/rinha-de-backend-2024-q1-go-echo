package customer

import (
	"context"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/database"
	"github.com/jackc/pgx/v5"
)

func GetCustomer(customerID int) (*Customer, error) {
	rows, err := database.Conn.Query(
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
