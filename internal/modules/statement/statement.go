package statement

import (
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/transaction"
)

type StatementSaldo struct {
	Total       int       `json:"total"`
	RequestDate time.Time `json:"data_extrato"`
	Limit      int       `json:"limite"`
}

type StatementResponse struct {
	Balance             StatementSaldo            `json:"saldo"`
	TransactionHistory []transaction.Transaction `json:"ultimas_transacoes"`
}
