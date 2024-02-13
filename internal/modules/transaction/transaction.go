package transaction

import "time"

type Transaction struct {
	Valor       int       `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type CreateTransactionRequest struct {
	Valor     int    `json:"valor" validate:"required,min=1"`
	Tipo      string `json:"tipo" validate:"required"`
	Descricao string `json:"descricao" validate:"required,min=1,max=10"`
}

type CreateTransactionResponse struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}
