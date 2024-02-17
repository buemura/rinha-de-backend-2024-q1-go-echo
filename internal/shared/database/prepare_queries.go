package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func PrepareQueries(conn *pgx.Conn) {
    _, err := conn.Prepare(context.Background(), "get_customer", "SELECT id, account_limit, account_balance FROM customers WHERE id = $1")
    if err != nil {
        log.Fatalf("Failed to prepare query 'get_customer': %v", err)
    }

	UpdateCustStmt, err = conn.Prepare(context.Background(), "update_balance", "UPDATE customers SET account_balance = $1 WHERE id = $2")
	if err != nil {
        log.Fatalf("Failed to prepare query 'update_balance': %v", err)
    }

	InsertTrxStmt, err = conn.Prepare(context.Background(), "insert_transaction", "INSERT INTO transactions (customer_id, amount, type, description, created_at) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
        log.Fatalf("Failed to prepare query 'insert_transaction': %v", err)
    }

    fmt.Println("Queries preparadas com sucesso!")
}