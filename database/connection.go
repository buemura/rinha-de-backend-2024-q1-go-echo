package database

import (
	"context"
	"fmt"
	"os"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/config"
	"github.com/jackc/pgx/v5"
)

var (
	Conn *pgx.Conn
)

func Connect() {
	conn, err := pgx.Connect(context.Background(), config.DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Conn = conn
}
