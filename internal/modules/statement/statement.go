package statement

import (
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/transaction"
)

type StatementSaldo struct {
	Total       int       `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      int       `json:"limite"`
}

type StatementResponse struct {
	Saldo             StatementSaldo            `json:"saldo"`
	UltimasTransacoes []transaction.Transaction `json:"ultimas_transacoes"`
}
