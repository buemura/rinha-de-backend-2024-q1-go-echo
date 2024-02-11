package customer

import "errors"

var ErrCustomerNotFound = errors.New("customer not found")

var ErrCustomerNoLimit = errors.New("customer has no limit")

type CustomerBalance struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Limite int    `json:"limite"`
	Saldo  int    `json:"saldo"`
}
