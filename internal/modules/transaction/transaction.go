package transaction

import "time"

type Transaction struct {
	Amount       int       `json:"valor"`
	Type         string    `json:"tipo"`
	Description  string    `json:"descricao"`
	CreatedAt    time.Time `json:"realizada_em"`
}

type CreateTransactionRequest struct {
	Amount      int    `json:"valor" validate:"required,min=1"`
	Type        string `json:"tipo" validate:"required"`
	Description string `json:"descricao" validate:"required,min=1,max=10"`
}

type CreateTransactionResponse struct {
	Limit    int `json:"limite"`
	Balance  int `json:"saldo"`
}
