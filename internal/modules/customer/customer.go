package customer

import "errors"

var ErrCustomerNotFound = errors.New("customer not found")

var ErrCustomerNoLimit = errors.New("customer has no limit")

type Customer struct {
	ID 				int
	AccountLimit    int `json:"limite"`
	AccountBalance  int `json:"saldo"`
}
