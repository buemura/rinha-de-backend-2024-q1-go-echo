package transaction

import (
	"context"
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/database"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/customer"
	"github.com/jackc/pgx/v5"
)

func GetTransactions(customerID int) ([]*Transaction, error) {
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

	transactions, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[Transaction])
	if err != nil {
		return nil, err
	}
	if transactions == nil {
		transactions = []*Transaction{}
	}

	return transactions, nil
}

func CreateTransaction(customerID int, data *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	customerBalance, err := customer.GetCustomerBalance(customerID)
	if err != nil {
		if customerBalance == nil {
			return nil, customer.ErrCustomerNotFound
		}
		return nil, err
	}

	if data.Tipo == "c" {
		customerBalance.Saldo += data.Valor
	}

	if data.Tipo == "d" {
		if (customerBalance.Saldo-data.Valor)*-1 > customerBalance.Limite {
			return nil, customer.ErrCustomerNoLimit
		}
		customerBalance.Saldo -= data.Valor
	}

	err = insertTransaction(customerID, customerBalance.Saldo, data.Valor, string(data.Tipo), data.Descricao)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionResponse{
		Saldo:  customerBalance.Saldo,
		Limite: customerBalance.Limite,
	}, nil
}

func insertTransaction(customerID, balance, trxAmount int, trxType, description string) error {
	tx, err := database.Conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	batch := &pgx.Batch{}
	batch.Queue(`
		INSERT INTO transacoes (cliente_id, valor, tipo, descricao, realizada_em) 
		VALUES ($1, $2, $3, $4, $5)
	`, customerID, trxAmount, trxType, description, time.Now())

	batch.Queue(`
		UPDATE saldos SET valor = $1 WHERE cliente_id = $2
	`, balance, customerID)

	bRes := tx.SendBatch(context.Background(), batch)
	_, err = bRes.Exec()
	if err != nil {
		return err
	}
	err = bRes.Close()
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
