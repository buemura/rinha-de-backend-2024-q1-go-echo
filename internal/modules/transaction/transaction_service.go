package transaction

import (
	"context"
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/customer"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/database"
	"github.com/jackc/pgx/v5"
)

func GetTransactions(customerID int) ([]Transaction, error) {
	rows, err := database.Conn.Query(
		context.Background(),
		`
        SELECT amount, type, description, created_at
        FROM transactions
        WHERE customer_id = $1
        ORDER BY created_at DESC
        LIMIT 10
        `,
		customerID,
	)
	if err != nil {
		return nil, err
	}

	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Transaction])
	if err != nil {
		return nil, err
	}
	if transactions == nil {
		transactions = []Transaction{}
	}

	return transactions, nil
}

func CreateTransaction(customerID int, trx *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	cust, err := customer.GetCustomer(customerID)
	if err != nil {
		return nil, err
	}
	if cust == nil {
		return nil, customer.ErrCustomerNotFound
	}
	trxRes, err := insertTransaction(cust, trx)
	if err != nil {
		return nil, err
	}
	return trxRes, nil
}

func insertTransaction(cust *customer.Customer, trx *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	tx, err := database.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	var balance int
	err = tx.QueryRow(context.Background(), "SELECT account_balance FROM customers WHERE id = $1 FOR UPDATE", cust.ID).Scan(&balance)
	if err != nil {
		return nil, err
	}

	if trx.Type == "c" {
		balance += trx.Amount
	}
	if trx.Type == "d" {
		balance -= trx.Amount
	}
	if cust.AccountLimit + balance < 0 {
		return nil, customer.ErrCustomerNoLimit
	}

	batch := &pgx.Batch{}
	batch.Queue(`
        INSERT INTO transactions (customer_id, amount, type, description, created_at) 
        VALUES ($1, $2, $3, $4, $5)
    `, cust.ID, trx.Amount, trx.Type, trx.Description, time.Now())
	batch.Queue(`
        UPDATE customers SET account_balance = $1 WHERE id = $2
    `, balance, cust.ID)

	bRes := tx.SendBatch(context.Background(), batch)
	if _, err := bRes.Exec(); err != nil {
		return nil, err
	}
	if err := bRes.Close(); err != nil {
		return nil, err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}


	return &CreateTransactionResponse{
		Balance:  balance,
		Limit: 	  cust.AccountLimit,
	}, nil
}
