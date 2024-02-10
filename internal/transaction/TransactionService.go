package transaction

import (
	"context"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/database"
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
