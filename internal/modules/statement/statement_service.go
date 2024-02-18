package statement

import (
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/customer"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/transaction"
)

func GetStatement(customerID int) (*StatementResponse, error) {
	customerBalance, err := customer.GetCustomer(customerID)
	if err != nil {
		if customerBalance == nil {
			return nil, customer.ErrCustomerNotFound
		}
		return nil, err
	}

	transactions, err := transaction.GetTransactions(customerID)
	if err != nil {
		return nil, err
	}

	return &StatementResponse{
		Saldo: StatementSaldo{
			Total:       customerBalance.AccountBalance,
			Limite:      customerBalance.AccountLimit,
			DataExtrato: time.Now(),
		},
		UltimasTransacoes: transactions,
	}, nil
}
