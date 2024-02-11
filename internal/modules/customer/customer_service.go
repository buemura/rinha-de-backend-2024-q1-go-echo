package customer

import (
	"context"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/database"
	"github.com/jackc/pgx/v5"
)

func GetCustomerBalance(customerID int) (*CustomerBalance, error) {
	rows, err := database.Conn.Query(
		context.Background(),
		"SELECT c.id, c.nome, c.limite, s.valor AS saldo FROM clientes c INNER JOIN saldos s ON c.id = s.cliente_id WHERE c.id = $1",
		customerID,
	)
	if err != nil {
		return nil, err
	}

	customerBalance, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[CustomerBalance])
	if err != nil {
		return nil, err
	}
	return customerBalance, nil
}
