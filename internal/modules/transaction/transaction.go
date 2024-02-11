package transaction

import "time"

type TransactionType string

const (
	Credit TransactionType = "c"
	Debit  TransactionType = "d"
)

type Transaction struct {
	Valor       int             `json:"valor"`
	Tipo        TransactionType `json:"tipo"`
	Descricao   string          `json:"descricao"`
	RealizadaEm time.Time       `json:"realizada_em"`
}

type CreateTransactionRequest struct {
	Valor     int             `json:"valor" validate:"required,min=1"`
	Tipo      TransactionType `json:"tipo" validate:"required"`
	Descricao string          `json:"descricao" validate:"required,min=1,max=10"`
}

type CreateTransactionResponse struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}
