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
	trxRes, err := insertTransaction(customerID, trx)
	if err != nil {
		return nil, err
	}
	return trxRes, nil
}

func insertTransaction(customerID int, trx *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	tx, err := database.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	var id, accountBalance, accountLimit int
	err = tx.QueryRow(
		context.Background(), 
		"SELECT id, account_limit, account_balance FROM customers WHERE id = $1 FOR UPDATE", 
		customerID,
	).Scan(&id, &accountLimit, &accountBalance)
	if err != nil {
		if id == 0 {
			return nil, customer.ErrCustomerNotFound
		}
		return nil, err
	}

	if trx.Type == "c" {
		accountBalance += trx.Amount
	}
	if trx.Type == "d" {
		accountBalance -= trx.Amount
	}
	if accountLimit + accountBalance < 0 {
		return nil, customer.ErrCustomerNoLimit
	}

	batch := &pgx.Batch{}
	batch.Queue(`
        INSERT INTO transactions (customer_id, amount, type, description, created_at) 
        VALUES ($1, $2, $3, $4, $5)
    `, customerID, trx.Amount, trx.Type, trx.Description, time.Now())
	batch.Queue(`
        UPDATE customers SET account_balance = $1 WHERE id = $2
    `, accountBalance, customerID)

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
		Balance:  accountBalance,
		Limit: 	  accountLimit,
	}, nil
}
