package statement

import (
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/customer"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/transaction"
)

type StatementService struct {
	cService customer.CustomerService
	tService transaction.TransactionService
}

func NewStatementService(cService customer.CustomerService, tService transaction.TransactionService) *StatementService {
	return &StatementService{
		cService: cService,
		tService: tService,
	}
}

func (s *StatementService) GetStatement(customerID int) (*StatementResponse, error) {
	customerBalance, err := s.cService.GetCustomer(customerID)
	if err != nil {
		if customerBalance == nil {
			return nil, customer.ErrCustomerNotFound
		}
		return nil, err
	}

	transactions, err := s.tService.GetTransactions(customerID)
	if err != nil {
		return nil, err
	}

	return &StatementResponse{
		Balance: StatementSaldo{
			Total:       customerBalance.AccountBalance,
			Limit:      customerBalance.AccountLimit,
			RequestDate: time.Now(),
		},
		TransactionHistory: transactions,
	}, nil
}
