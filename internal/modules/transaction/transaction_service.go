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
        SELECT valor, tipo, descricao, realizada_em
        FROM transacoes
        WHERE cliente_id = $1
        order by realizada_em desc
        limit 10
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

func CreateTransaction(customerID int, trx *CreateTransactionRequest) (CreateTransactionResponse, error) {
	customerBalance, err := customer.GetCustomerBalance(customerID)
	if err != nil {
		if customerBalance == (customer.CustomerBalance{}) {
			return CreateTransactionResponse{}, customer.ErrCustomerNotFound
		}
		return CreateTransactionResponse{}, err
	}

	trxRes, err := insertTransaction(customerID, customerBalance.Limite, trx.Valor, string(trx.Tipo), trx.Descricao)
	if err != nil {
		return CreateTransactionResponse{}, err
	}
	return trxRes, nil
}

func insertTransaction(customerID, limit, trxAmount int, trxType, description string) (CreateTransactionResponse, error) {
	tx, err := database.Conn.Begin(context.Background())
	if err != nil {
		return CreateTransactionResponse{}, err
	}
	defer tx.Rollback(context.Background())

	var balance int
	err = tx.QueryRow(context.Background(), "SELECT valor FROM saldos WHERE cliente_id = $1 FOR UPDATE", customerID).Scan(&balance)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	if trxType == "c" {
		balance += trxAmount
	}
	if trxType == "d" {
		if (balance-trxAmount)*-1 > limit {
			return CreateTransactionResponse{}, customer.ErrCustomerNoLimit
		}
		balance -= trxAmount
	}

	batch := &pgx.Batch{}
	batch.Queue(`
        INSERT INTO transacoes (cliente_id, valor, tipo, descricao, realizada_em) 
        VALUES ($1, $2, $3, $4, $5)
    `, customerID, trxAmount, trxType, description, time.Now())
	batch.Queue(`
        UPDATE saldos SET valor = $1 WHERE cliente_id = $2
    `, balance, customerID)

	bRes := tx.SendBatch(context.Background(), batch)
	if _, err := bRes.Exec(); err != nil {
		return CreateTransactionResponse{}, err
	}
	if err := bRes.Close(); err != nil {
		return CreateTransactionResponse{}, err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return CreateTransactionResponse{}, err
	}

	return CreateTransactionResponse{
		Saldo:  balance,
		Limite: limit,
	}, nil
}
