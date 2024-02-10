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
